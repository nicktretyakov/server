package booking

import (
	"context"

	"be/internal/acl"
	"be/internal/model"
)

func (s Service) UpdateBooking(
	ctx context.Context,
	user model.User,
	updatedBooking *model.Booking,
	portalCodeSupervisor, portalCodeAssignee uint64,
) (*model.Booking, error) {
	if acl.IsUserWithoutPrivileges(user, *updatedBooking) {
		return nil, ErrPermissionDenied
	}

	if !acl.CanUpdateAddress(user, updatedBooking) {
		return nil, ErrPermissionDenied
	}

	if !changeStatusInitialAuthor(user, *updatedBooking) {
		updatedBooking.Status = model.OnRegisterAddressStatus
	}

	updatedBooking, err := s.SetBookingAssignee(ctx, updatedBooking, portalCodeAssignee)
	if err != nil {
		return nil, err
	}

	updatedBooking, err = s.SetBookingSupervisor(ctx, updatedBooking, portalCodeSupervisor)
	if err != nil {
		return nil, err
	}

	if acl.IsHeadOfBooking(user) && updatedBooking.Status == model.ConfirmedAddressStatus {
		if err = s.reportService.CheckReports(ctx, &updatedBooking.ID, updatedBooking.Timeline.Periods()); err != nil {
			return nil, err
		}
	}

	return s.store.Booking().Update(ctx, *updatedBooking)
}

func changeStatusInitialAuthor(user model.User, booking model.Booking) bool {
	return booking.IsAuthor(user) && booking.Status.Eq(model.InitialAddressStatus)
}
