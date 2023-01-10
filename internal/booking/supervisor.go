package booking

import (
	"context"

	"be/internal/model"
)

func (s Service) SetBookingSupervisor(ctx context.Context, book *model.Booking, portalCode uint64) (*model.Booking, error) {
	supervisor, err := s.userService.GetOrCreateUserByPortalCode(ctx, portalCode)
	if err != nil {
		return nil, err
	}

	if book.Supervisor != nil {
		if err = s.store.Booking().DeleteBookingUser(ctx, model.BookingUser{
			BookingID: book.ID,
			UserID:    book.Supervisor.ID,
			Role:      model.SupervisorBookingUserRole,
		}); err != nil {
			return nil, err
		}
	}

	book.Supervisor = supervisor

	if err = s.store.Booking().CreateBookingUser(ctx, model.BookingUser{
		BookingID: book.ID,
		UserID:    supervisor.ID,
		Role:      model.SupervisorBookingUserRole,
	}); err != nil {
		return nil, err
	}

	return book, nil
}
