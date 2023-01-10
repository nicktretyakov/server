package dbmodel

import (
	"github.com/google/uuid"

	"be/internal/model"
)

type BookingUser struct {
	BookingID uuid.UUID             `db:"booking_id"`
	Role      model.BookingUserRole `db:"role"`
	User      User                  `db:"user"`
	RoleTitle *string               `db:"role_description"`
}
