package outmember

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"be/internal/booking"
	addressStore "be/internal/datastore/address"
	"be/internal/datastore/filters"
	"be/internal/datastore/outmember"
	"be/internal/model"
	"be/internal/room"
)

type Address struct {
	ID          uuid.UUID
	typeAddress addressStore.TypeAddress
	status      model.Status
	period      []model.Period
}

func (s Service) newOutmember(
	ctx context.Context,
	user model.User,
	typeAddress addressStore.TypeAddress,
	outmember IOutmember,
) error {
	var (
		book    *model.Booking
		roo    *model.Room
		address Address
		err     error
	)

	switch typeAddress {
	case addressStore.BookingAddressType:
		book, err = s.store.Booking().FindByID(ctx, outmember.AddressID())
		if err != nil {
			return booking.ErrBookingNotFound
		}

		address.ID = book.ID
		address.status = book.Status
		address.typeAddress = typeAddress
		address.period = book.Timeline.Periods()

		if !outmember.ACL(user, book) {
			return ErrPermissionDenied
		}

	case addressStore.RoomAddressType:
		roo, err = s.store.Room().FindByID(ctx, outmember.AddressID())
		if err != nil {
			return room.ErrRoomNotFound
		}

		address.ID = roo.ID
		address.status = roo.Status
		address.typeAddress = typeAddress
		address.period = nil

		if !outmember.ACL(user, roo) {
			return ErrPermissionDenied
		}

	case addressStore.UnknownAddressType:
		return errors.New("unknown invest object type")
	default:
		return errors.New("unknown invest object type")
	}

	if _, err = s.store.Outmember().Create(ctx, outmember.Source(), typeAddress); err != nil {
		return err
	}

	return s.outmemberCreated(ctx, &address, outmember)
}

//nolint:gocognit
func (s Service) outmemberCreated(ctx context.Context, address *Address, agr IOutmember) (err error) {
	switch {
	case address.status == model.OnRegisterAddressStatus && agr.Type() == model.InitialOutmemberType:
		err = s.onRegisterAddressOutmemberCreated(ctx, address)
	case address.status == model.FinalizeOnRegisterStatus && agr.Type() == model.FinalRegisterOutmemberType:
		err = s.confirmedBookingFinalRegisterOutmemberCreated(ctx, address)
	case address.status == model.FinalizeOnAgreeStatus && agr.Type() == model.FinalApprovalOutmemberType:
		err = s.confirmedBookingFinalApprovalOutmemberCreated(ctx, address)
	case address.status == model.OnAgreeAddressStatus && agr.Type() == model.ApprovalOutmemberType:
		err = s.onAgreeAddressOutmemberCreated(ctx, address)
	}

	return err
}

func (s Service) onRegisterAddressOutmemberCreated(ctx context.Context, address *Address) error {
	var (
		baseFilters = []filters.Filter{
			outmember.NewOutmemberTypeFilter(model.InitialOutmemberType),
			outmember.NewAddressIDFilter(address.ID),
		}
		initialOutmembers []model.Outmember
		err               error
	)

	if initialOutmembers, err = s.store.Outmember().List(ctx, address.typeAddress, baseFilters...); err != nil {
		return err
	}

	if len(initialOutmembers) > 0 && initialOutmembers[0].Result == model.AcceptOutmemberResult {
		return s.updateStatus(ctx, address.typeAddress, address.ID, model.OnAgreeAddressStatus)
	}

	if len(initialOutmembers) > 0 && initialOutmembers[0].Result == model.DeclineOutmemberResult {
		return s.updateStatus(ctx, address.typeAddress, address.ID, model.DeclinedAddressStatus)
	}

	return nil
}

func (s Service) confirmedBookingFinalRegisterOutmemberCreated(ctx context.Context, address *Address) error {
	var (
		baseFilters = []filters.Filter{
			outmember.NewOutmemberTypeFilter(model.FinalRegisterOutmemberType),
			outmember.NewAddressIDFilter(address.ID),
		}
		finalOutmembers []model.Outmember
		err             error
	)

	if finalOutmembers, err = s.store.Outmember().List(ctx, address.typeAddress, baseFilters...); err != nil {
		return err
	}

	if len(finalOutmembers) == 0 {
		return nil
	}

	switch finalOutmembers[0].Result {
	case model.DeclineOutmemberResult:
		return s.declineFinalizeAddress(ctx, address)
	case model.AcceptOutmemberResult:
		if err = s.store.Booking().UpdateStatus(ctx, address.ID, model.FinalizeOnAgreeStatus); err != nil {
			return err
		}

		return s.store.FinalReport().UpdateStatus(ctx, address.ID, model.OnAgreeFinalReportStatus)
	}

	return err
}

func (s Service) confirmedBookingFinalApprovalOutmemberCreated(ctx context.Context, address *Address) error {
	var (
		baseFilters = []filters.Filter{
			outmember.NewOutmemberTypeFilter(model.FinalApprovalOutmemberType),
			outmember.NewAddressIDFilter(address.ID),
		}
		bookingFinalOutmembers []model.Outmember
		err                    error
	)

	if bookingFinalOutmembers, err = s.store.Outmember().List(ctx, address.typeAddress, baseFilters...); err != nil {
		return err
	}

	if len(bookingFinalOutmembers) == 0 {
		return nil
	}

	switch bookingFinalOutmembers[0].Result {
	case model.DeclineOutmemberResult:
		if err = s.declineFinalizeAddress(ctx, address); err != nil {
			return err
		}

		return s.store.Outmember().BulkDelete(ctx, address.ID, model.FinalRegisterOutmemberType, address.typeAddress)
	case model.AcceptOutmemberResult:
		if err := s.store.FinalReport().UpdateStatus(ctx, address.ID, model.ConfirmedFinalReportStatus); err != nil {
			return err
		}

		return s.store.Booking().UpdateStatus(ctx, address.ID, model.DoneAddressStatus)
	}

	return err
}

func (s Service) onAgreeAddressOutmemberCreated(ctx context.Context, address *Address) error {
	var (
		baseFilters = []filters.Filter{
			outmember.NewOutmemberTypeFilter(model.ApprovalOutmemberType),
			outmember.NewAddressIDFilter(address.ID),
		}
		bookingOutmembers []model.Outmember
		err               error
	)

	if bookingOutmembers, err = s.store.Outmember().List(ctx, address.typeAddress, baseFilters...); err != nil {
		return err
	}

	if len(bookingOutmembers) > 0 && bookingOutmembers[0].Result == model.AcceptOutmemberResult {
		return s.confirmAddress(ctx, address)
	}

	if len(bookingOutmembers) > 0 && bookingOutmembers[0].Result == model.DeclineOutmemberResult {
		return s.declineAddress(ctx, address)
	}

	return nil
}

func (s Service) confirmAddress(ctx context.Context, address *Address) error {
	if err := s.updateStatus(ctx, address.typeAddress, address.ID, model.ConfirmedAddressStatus); err != nil {
		return err
	}

	if address.typeAddress == addressStore.RoomAddressType {
		return nil
	}

	return s.reportService.CheckReports(ctx, &address.ID, address.period)
}

func (s Service) declineAddress(ctx context.Context, address *Address) error {
	if err := s.updateStatus(ctx, address.typeAddress, address.ID, model.DeclinedAddressStatus); err != nil {
		return err
	}

	return s.store.Outmember().BulkDelete(ctx, address.ID, model.InitialOutmemberType, address.typeAddress)
}

func (s Service) declineFinalizeAddress(ctx context.Context, address *Address) error {
	if err := s.store.Booking().UpdateStatus(ctx, address.ID, model.FinalizeReportDeclined); err != nil {
		return err
	}

	return s.store.FinalReport().UpdateStatus(ctx, address.ID, model.DeclinedFinalReportStatus)
}

func (s Service) updateStatus(
	ctx context.Context,
	typeAddress addressStore.TypeAddress,
	addressUUID uuid.UUID,
	status model.Status,
) error {
	switch typeAddress {
	case addressStore.BookingAddressType:
		return s.store.Booking().UpdateStatus(ctx, addressUUID, status)
	case addressStore.RoomAddressType:
		return s.store.Room().UpdateStatus(ctx, addressUUID, status)
	case addressStore.UnknownAddressType:
		return errors.New("unknown  type")
	default:
		return errors.New("unknown  type")
	}
}
