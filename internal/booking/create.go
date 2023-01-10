package booking

import (
	"context"

	"be/internal/model"
)

func (s Service) CreateBooking(ctx context.Context, user model.User, initialBooking model.Booking,
	portalCode uint64,
) (*model.Booking, error) {
	supervisor, err := s.userService.GetOrCreateUserByPortalCode(ctx, portalCode)
	if err != nil {
		return nil, err
	}

	initialBooking.Supervisor = supervisor
	initialBooking.Author = &user

	return s.store.Booking().Create(ctx, initialBooking)
}
