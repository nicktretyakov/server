package room

import (
	"context"

	"github.com/google/uuid"
	bookingpb "be/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"be/internal/acl"
	"be/internal/datastore/filters"
	"be/internal/datastore/room"
	"be/internal/datastore/sorting"
	"be/internal/model"
	"be/pkg/auth"
	"be/pkg/server/pbs"
)

func (s Service) GetRoom(ctx context.Context, req *bookingpb.GetRoomRequest) (*bookingpb.GetRoomResponse, error) {
	roomID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid room_id")
	}

	room, err := s.roomFromDB(ctx, roomID)
	if err != nil {
		return nil, err
	}

	return &bookingpb.GetRoomResponse{Room: room}, nil
}

func (s Service) GetArchivedRooms(ctx context.Context, req *bookingpb.GetRoomsRequest) (*bookingpb.GetRoomsResponse, error) {
	user := auth.FromContext(ctx)
	if user.Role != model.Admin {
		return nil, status.New(codes.PermissionDenied, "head of booking booking only").Err()
	}

	rooms, total, err := s.getRoomByFilter(ctx, []model.State{model.ArchivedAddressState}, req)
	if err != nil {
		return nil, err
	}

	return &bookingpb.GetRoomsResponse{
		Rooms: pbs.PbRooms(rooms),
		Count:    uint32(total),
	}, nil
}

func (s Service) GetRooms(ctx context.Context, req *bookingpb.GetRoomsRequest) (*bookingpb.GetRoomsResponse, error) {
	rooms, total, err := s.getRoomByFilter(ctx, []model.State{model.PublishedAddressState}, req)
	if err != nil {
		return nil, err
	}

	return &bookingpb.GetRoomsResponse{
		Rooms: pbs.PbRooms(rooms),
		Count:    uint32(total),
	}, nil
}

func (s Service) getRoomByFilter(
	ctx context.Context,
	states []model.State,
	req *bookingpb.GetRoomsRequest,
) ([]model.Room, uint64, error) {
	var (
		offset       uint64 = 0
		limit        uint64 = 50
		queryFilters []filters.Filter
		err          error
	)

	if req.GetOffset() > 0 {
		offset = uint64(req.GetOffset())
	}

	if req.GetLimit() > 0 {
		limit = uint64(req.GetLimit())
	}

	user := auth.FromContext(ctx)

	queryFilters, err = s.withFilter(ctx, user, states, req.GetFilter())
	if err != nil {
		return nil, 0, err
	}

	if req.GetQuery() != "" {
		queryFilters = append(queryFilters, room.NewSearchFilter(req.GetQuery()))
	}

	sort := withSorting(req.Sorting)

	return s.store.Room().ActiveList(ctx, limit, offset, sort, queryFilters...)
}

func (s Service) withFilter(
	ctx context.Context,
	user *model.User,
	states []model.State,
	pb *bookingpb.RoomFilter,
) ([]filters.Filter, error) {
	filtersToApply := make([]filters.Filter, 0)

	list := []struct {
		filterCall func() (filters.Filter, error)
	}{
		{
			filterCall: func() (filters.Filter, error) { return authorFilter(pb, user.ID) },
		},
		{
			filterCall: func() (filters.Filter, error) { return confirmationAwaitFilter(pb, user) },
		},
		{
			filterCall: func() (filters.Filter, error) { return getCreatedAt(pb) },
		},
		{
			filterCall: func() (filters.Filter, error) { return getCreationDate(pb) },
		},
		{
			filterCall: func() (filters.Filter, error) { return getStatuses(pb) },
		},
		{
			filterCall: func() (filters.Filter, error) { return s.getOwners(ctx, pb) },
		},
		{
			filterCall: func() (filters.Filter, error) { return s.getEmployees(ctx, pb) },
		},
		{
			filterCall: func() (filters.Filter, error) { return getStateFilter(states) },
		},
	}

	for _, v := range list {
		filters, err := v.filterCall()
		if err != nil {
			return nil, err
		}

		if filters != nil {
			filtersToApply = append(filtersToApply, filters)
		}
	}

	return filtersToApply, nil
}

func authorFilter(pb *bookingpb.RoomFilter, userID uuid.UUID) (filters.Filter, error) {
	if pb == nil || !pb.GetCreatedByMe() {
		return nil, nil
	}

	return room.NewAuthorFilter(userID), nil
}

func confirmationAwaitFilter(pb *bookingpb.RoomFilter, user *model.User) (filters.Filter, error) {
	if pb == nil || !pb.GetAwaitsMe() {
		return nil, nil
	}

	if acl.IsHeadOfBooking(*user) {
		return room.NewPOHeadConfirmationWaitFilter(), nil
	}

	return room.NewAssigneeConfirmationWaitFilter(user.ID), nil
}

func getCreatedAt(pb *bookingpb.RoomFilter) (filters.Filter, error) {
	if pb == nil || pb.GetCreatedAt() == nil {
		return nil, nil
	}

	return room.NewCreatedAtFilter(pb.GetCreatedAt().GetStart(), pb.GetCreatedAt().GetEnd())
}

func getCreationDate(pb *bookingpb.RoomFilter) (filters.Filter, error) {
	if pb == nil || pb.GetCreationDate() == nil {
		return nil, nil
	}

	return room.NewCreationDateFilter(pb.GetCreationDate().GetStart(), pb.GetCreationDate().GetEnd())
}

func getStatuses(pb *bookingpb.RoomFilter) (filters.Filter, error) {
	if pb == nil || pb.GetStatus() == nil {
		return nil, nil
	}

	statuses := make([]model.Status, 0, len(pb.GetStatus()))

	for _, st := range pb.GetStatus() {
		statuses = append(statuses, model.Status(st))
	}

	return room.NewStatusFilter(statuses), nil
}

func (s Service) getOwners(ctx context.Context, pb *bookingpb.RoomFilter) (filters.Filter, error) {
	if pb == nil || pb.GetPortalCodeOwners() == nil {
		return nil, nil
	}

	ownerIDs, err := s.getUsersByPortalCode(ctx, pb.GetPortalCodeOwners())
	if err != nil {
		return nil, err
	}

	return room.NewOwnerFilter(ownerIDs)
}

func (s Service) getEmployees(ctx context.Context, pb *bookingpb.RoomFilter) (filters.Filter, error) {
	if pb == nil || pb.GetPortalCodeEmployees() == nil {
		return nil, nil
	}

	employeeIDs, err := s.getUsersByPortalCode(ctx, pb.GetPortalCodeEmployees())
	if err != nil {
		return nil, err
	}

	return room.NewEmployeeFilter(employeeIDs)
}

func withSorting(pbSorting *bookingpb.RoomSorting) (s sorting.Sorting) {
	switch pbSorting.GetType() {
	case bookingpb.RoomSorting_BY_PUBLISH_DATE:
		s = room.NewSortingByPublishDate(pbSorting.GetAsc())
	case bookingpb.RoomSorting_BY_CREATION_DATE_DATE:
		s = room.NewSortingByCreationDateSorting(pbSorting.GetAsc())
	default:
		s = room.NewSortingByPublishDate(false)
	}

	return s
}

func (s Service) getUsersByPortalCode(ctx context.Context, portalCodes []uint32) ([]uuid.UUID, error) {
	pCodes := make([]uint64, 0, len(portalCodes))

	for _, portalCode := range portalCodes {
		pCodes = append(pCodes, uint64(portalCode))
	}

	usersID, err := s.store.User().FindUsersIDByPortalCode(ctx, pCodes)
	if err != nil {
		return nil, err
	}

	return usersID, nil
}

func getStateFilter(states []model.State) (filters.Filter, error) {
	return room.NewStateFilter(states), nil
}
