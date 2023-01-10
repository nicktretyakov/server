package booking

import (
	"context"

	"be/internal/model"
)

func (s Service) SetBookingAssignee(ctx context.Context, book *model.Booking, portalCode uint64) (*model.Booking, error) {
	assignedUser, err := s.userService.GetOrCreateUserByPortalCode(ctx, portalCode)
	if err != nil {
		return nil, err
	}

	if book.IsAssigned() {
		if err = s.store.Booking().DeleteBookingUser(ctx, model.BookingUser{
			BookingID: book.ID,
			UserID:    book.Assignee.ID,
			Role:      model.AssigneeBookingUserRole,
		}); err != nil {
			return nil, err
		}
	}

	book.Assignee = assignedUser

	if err = s.store.Booking().CreateBookingUser(ctx, model.BookingUser{
		BookingID: book.ID,
		UserID:    assignedUser.ID,
		Role:      model.AssigneeBookingUserRole,
	}); err != nil {
		return nil, err
	}

	return book, nil
}
