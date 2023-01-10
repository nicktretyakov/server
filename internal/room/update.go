package room

import (
	"context"

	"github.com/google/uuid"

	"be/internal/acl"
	"be/internal/model"
)

func (s Service) Update(ctx context.Context, user model.User, room model.Room, portalCodeOwner, portalCodeAssignee uint64,
	equipmentIDs, slotIDs, bookingIDs []uuid.UUID,
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

	if !acl.CanUpdateAddress(user, room) {
		return nil, ErrPermissionDenied
	}

	return s.store.Room().Update(ctx, room, equipmentIDs, slotIDs, bookingIDs)
}
