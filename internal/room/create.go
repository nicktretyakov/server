package room

import (
	"context"

	"github.com/google/uuid"

	"be/internal/model"
)

func (s Service) Create(ctx context.Context, room model.Room, portalCodeOwner, portalCodeAssignee uint64,
	equipmentIDs, slotIDs []uuid.UUID,
) (*model.Room, error) {
	ownerUser, err := s.userService.GetOrCreateUserByPortalCode(ctx, portalCodeOwner)
	if err != nil {
		return nil, err
	}

	assigneeUser, err := s.userService.GetOrCreateUserByPortalCode(ctx, portalCodeAssignee)
	if err != nil {
		return nil, err
	}

	room.Owner = ownerUser
	room.Employee = assigneeUser

	return s.store.Room().Create(ctx, room, equipmentIDs, slotIDs)
}
