package room

import (
	"context"

	"github.com/google/uuid"

	"be/internal/model"
)

type IRoomService interface {
	Create(ctx context.Context, room model.Room, portalCodeOwner, portalCodeAssignee uint64,
		equipmentIDs, slotIDs []uuid.UUID) (*model.Room, error)
	Update(ctx context.Context, user model.User, room model.Room, portalCodeOwner, portalCodeAssignee uint64,
		equipmentIDs, slotIDs, bookingIDs []uuid.UUID) (*model.Room, error)
}
