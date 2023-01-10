package booking

import (
	"context"
	"time"

	bookingpb "be/proto"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"be/internal/acl"
	bookingStore "be/internal/datastore/booking"
	"be/internal/datastore/filters"
	"be/internal/datastore/sorting"
	"be/internal/lib"
	"be/internal/model"
	"be/pkg/auth"
	"be/pkg/server/pbs"
)

func (s Service) GetBooking(ctx context.Context, req *bookingpb.GetBookingRequest) (*bookingpb.GetBookingResponse, error) {
	bookingID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid booking_id")
	}

	book, err := s.bookingFromDB(ctx, bookingID)
	if err != nil {
		return nil, err
	}

	return &bookingpb.GetBookingResponse{Booking: book}, nil
}

func (s Service) GetArchivedBookings(ctx context.Context, req *bookingpb.GetListRequest) (*bookingpb.GetListResponse, error) {
	user := auth.FromContext(ctx)
	if user.Role != model.Admin {
		return nil, status.New(codes.PermissionDenied, "head of booking booking only").Err()
	}

	bookings, total, err := s.getBookingsByFilter(ctx, []model.State{model.ArchivedAddressState}, req)
	if err != nil {
		return nil, err
	}

	return &bookingpb.GetListResponse{
		Bookings: pbs.PbBookings(bookings),
		Count:    uint32(total),
	}, nil
}

func (s Service) GetList(ctx context.Context, req *bookingpb.GetListRequest) (*bookingpb.GetListResponse, error) {
	bookings, total, err := s.getBookingsByFilter(ctx, []model.State{model.PublishedAddressState}, req)
	if err != nil {
		return nil, err
	}

	return &bookingpb.GetListResponse{
		Bookings: pbs.PbBookings(bookings),
		Count:    uint32(total),
	}, nil
}

func (s Service) getBookingsByFilter(
	ctx context.Context,
	states []model.State,
	req *bookingpb.GetListRequest,
) ([]model.Booking, uint64, error) {
	var (
		offset  uint64 = 0
		limit   uint64 = 50
		user           = auth.FromContext(ctx)
		sorting sorting.Sorting
	)

	if req.GetOffset() > 0 {
		offset = uint64(req.GetOffset())
	}

	if req.GetLimit() > 0 {
		limit = uint64(req.GetLimit())
	}

	var queryFilters []filters.Filter
	if req.GetQuery() != "" {
		queryFilters = append(queryFilters, bookingStore.NewSearchFilter(req.GetQuery()))
	}

	f, err := s.withFilters(ctx, *user, states, req.Filter)
	if err != nil {
		return nil, 0, err
	}

	queryFilters = append(queryFilters, f...)

	sorting = withSorting(req.Sorting)

	return s.store.Booking().ActiveList(ctx, limit, offset, sorting, queryFilters...)
}

func (s Service) withFilters(
	ctx context.Context,
	user model.User,
	states []model.State,
	pbFilter *bookingpb.Filter,
) ([]filters.Filter, error) {
	filtersToApply := make([]filters.Filter, 0)

	list := []struct {
		filterCall func() (filters.Filter, error)
	}{
		{
			filterCall: func() (filters.Filter, error) { return getStatusFilterNew(pbFilter) },
		},
		{
			filterCall: func() (filters.Filter, error) { return timelineFinish(pbFilter) },
		},
		{
			filterCall: func() (filters.Filter, error) { return supervisorFilter(pbFilter, user) },
		},
		{
			filterCall: func() (filters.Filter, error) { return confirmationAwaitFilter(pbFilter, user) },
		},
		{
			filterCall: func() (filters.Filter, error) { return s.supervisorsFilter(ctx, pbFilter) },
		},
		{
			filterCall: func() (filters.Filter, error) { return departmentsFilter(pbFilter) },
		},
		{
			filterCall: func() (filters.Filter, error) { return slotFilter(pbFilter) },
		},
		{
			filterCall: func() (filters.Filter, error) { return stateFilter(states) },
		},
	}

	for _, v := range list {
		filter, err := v.filterCall()
		if err != nil {
			return nil, err
		}

		if filter != nil {
			filtersToApply = append(filtersToApply, filter)
		}
	}

	return filtersToApply, nil
}

func getStatusFilterNew(pbFilter *bookingpb.Filter) (filters.Filter, error) {
	if pbFilter == nil || pbFilter.GetStatus() == nil {
		return nil, nil
	}

	statuses := make([]model.Status, 0, len(pbFilter.GetStatus()))

	for _, statusPB := range pbFilter.GetStatus() {
		statuses = append(statuses, getStatus(statusPB))
	}

	return bookingStore.NewStatusFilter(statuses), nil
}

//nolint:cyclop,gocyclo
func getStatus(statusPB bookingpb.AddressStatus) model.Status {
	switch statusPB {
	case bookingpb.AddressStatus_INITIAL:
		return model.InitialAddressStatus
	case bookingpb.AddressStatus_ON_REGISTRATION:
		return model.OnRegisterAddressStatus
	case bookingpb.AddressStatus_ON_AGREE:
		return model.OnAgreeAddressStatus
	case bookingpb.AddressStatus_CONFIRMED:
		return model.ConfirmedAddressStatus
	case bookingpb.AddressStatus_DONE:
		return model.DoneAddressStatus
	case bookingpb.AddressStatus_DECLINED:
		return model.DeclinedAddressStatus
	case bookingpb.AddressStatus_FINALIZE_ON_REGISTRATION:
		return model.FinalizeOnRegisterStatus
	case bookingpb.AddressStatus_FINALIZE_ON_AGREE:
		return model.FinalizeOnAgreeStatus
	case bookingpb.AddressStatus_FINALIZE_DECLINED:
		return model.FinalizeReportDeclined
	case bookingpb.AddressStatus_ADDRESS_STATUS:
		return model.UnknownAddressStatus
	default:
		return model.UnknownAddressStatus
	}
}

func confirmationAwaitFilter(pbFilter *bookingpb.Filter, user model.User) (filters.Filter, error) {
	if pbFilter == nil || !pbFilter.GetAwaitsMe() {
		return nil, nil
	}

	if acl.IsHeadOfBooking(user) {
		return bookingStore.NewPOHeadConfirmationWaitFilter(), nil
	}

	return bookingStore.NewAssigneeConfirmationWaitFilter(user.ID), nil
}

func timelineFinish(pbFilter *bookingpb.Filter) (filters.Filter, error) {
	if pbFilter == nil || pbFilter.GetTimelineFinish() == nil {
		return nil, nil
	}

	endAt, err := time.ParseInLocation(pbs.TimeLayout, pbFilter.TimelineFinish.End, time.UTC)
	if err != nil {
		return nil, err
	}

	return bookingStore.NewTimeLineFilter(endAt), nil
}

func (s Service) supervisorsFilter(ctx context.Context, pbFilter *bookingpb.Filter) (filters.Filter, error) {
	if pbFilter == nil || pbFilter.GetPortalCodeSupervisors() == nil {
		return nil, nil
	}

	pCodes := make([]uint64, 0, len(pbFilter.GetPortalCodeSupervisors()))

	for _, portalCode := range pbFilter.GetPortalCodeSupervisors() {
		pCodes = append(pCodes, uint64(portalCode))
	}

	supervisorUUIDs, err := s.store.User().FindUsersIDByPortalCode(ctx, pCodes)
	if err != nil {
		return nil, err
	}

	return bookingStore.NewSupervisorManyFilter(supervisorUUIDs), nil
}

func supervisorFilter(pbFilter *bookingpb.Filter, user model.User) (filters.Filter, error) {
	if pbFilter == nil || !pbFilter.GetCreatedByMe() {
		return nil, nil
	}

	return bookingStore.NewSupervisorFilter(user.ID), nil
}

func departmentsFilter(pbFilter *bookingpb.Filter) (filters.Filter, error) {
	if pbFilter == nil || pbFilter.GetDepartment() == nil {
		return nil, nil
	}

	list := make([]model.DepartmentID, 0)

	for _, department := range pbFilter.Department {
		list = append(list, model.DepartmentID{
			ID: uuid.MustParse(department.Id),
		})
	}

	return bookingStore.NewDepartmentManyFilter(list), nil
}

func slotFilter(pbFilter *bookingpb.Filter) (filters.Filter, error) {
	if pbFilter == nil || pbFilter.GetSlot() == nil {
		return nil, nil
	}

	var from *model.Notification
	if pbFilter.Slot.SlotFrom != nil {
		from = lib.Notification(model.NewNotificationUnitsAndFragments(pbFilter.Slot.SlotFrom.GetUnits(), pbFilter.Slot.SlotFrom.GetFragments()))
	}

	var to *model.Notification
	if pbFilter.Slot.SlotTo != nil {
		to = lib.Notification(model.NewNotificationUnitsAndFragments(pbFilter.Slot.SlotTo.GetUnits(), pbFilter.Slot.SlotTo.GetFragments()))
	}

	return bookingStore.NewSlotFilter(from, to), nil
}

func stateFilter(states []model.State) (filters.Filter, error) {
	return bookingStore.NewStateFilter(states), nil
}

func withSorting(pbSorting *bookingpb.Sorting) (s sorting.Sorting) {
	switch pbSorting.GetType() {
	case bookingpb.SortType_BY_PUBLISH_DATE:
		s = bookingStore.SortingByPublishDate(pbSorting.GetAsc())
	case bookingpb.SortType_BY_END_DATE:
		s = bookingStore.SortingByEndDate(pbSorting.GetAsc())
	case bookingpb.SortType_BY_SLOT:
		s = bookingStore.SortingBySlot(pbSorting.GetAsc())
	}

	return s
}
