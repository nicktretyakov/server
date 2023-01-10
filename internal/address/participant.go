package address

import (
	"context"
	"github.com/google/uuid"
	"be/internal/acl"
	"be/internal/booking"
	"be/internal/datastore/address"
	"be/internal/model"
	"be/internal/room"
)

func (s Service) AddParticipant(
	ctx context.Context,
	user model.User,
	role string,
	portalCodeParticipant uint64,
	addressID uuid.UUID,
	typeAddress address.TypeAddress,
) (*model.User, string, error) {
	userByPortalCode, err := s.userService.GetOrCreateUserByPortalCode(ctx, portalCodeParticipant)
	if err != nil {
		return nil, "", err
	}

	switch typeAddress {
	case address.RoomAddressType:
		if err = s.addParticipantRoom(ctx, user, *userByPortalCode, addressID, role); err != nil {
			return nil, "", err
		}
	case address.BookingAddressType:
		if err = s.addParticipantBooking(ctx, user, *userByPortalCode, addressID, role); err != nil {
			return nil, "", err
		}
	case address.UnknownAddressType:
		return nil, "", ErrUnknownTypeAddress
	default:
		return nil, "", ErrUnknownTypeAddress
	}

	return userByPortalCode, role, nil
}

func (s Service) UpdateParticipant(
	ctx context.Context,
	user model.User,
	role string,
	portalCodeParticipantBase, portalCodeParticipantChange uint64,
	addressID uuid.UUID,
	typeAddress address.TypeAddress,
) (*model.User, string, error) {
	userBase, err := s.userService.GetOrCreateUserByPortalCode(ctx, portalCodeParticipantBase)
	if err != nil {
		return nil, "", err
	}

	userChange, err := s.userService.GetOrCreateUserByPortalCode(ctx, portalCodeParticipantChange)
	if err != nil {
		return nil, "", err
	}

	switch typeAddress {
	case address.RoomAddressType:
		if err = s.updateParticipantRoom(ctx, user, *userBase, *userChange, addressID, role); err != nil {
			return nil, "", err
		}
	case address.BookingAddressType:
		if err = s.updateParticipantBooking(ctx, user, *userBase, *userChange, addressID, role); err != nil {
			return nil, "", err
		}
	case address.UnknownAddressType:
		return nil, "", ErrUnknownTypeAddress
	default:
		return nil, "", ErrUnknownTypeAddress
	}

	return userChange, role, nil
}

func (s Service) RemoveParticipant(
	ctx context.Context,
	user model.User,
	portalCodeParticipant uint64,
	addressID uuid.UUID,
	typeAddress address.TypeAddress,
) error {
	userByPortalCode, err := s.userService.GetOrCreateUserByPortalCode(ctx, portalCodeParticipant)
	if err != nil {
		return err
	}

	switch typeAddress {
	case address.RoomAddressType:
		if err = s.removeParticipantRoom(ctx, user, *userByPortalCode, addressID); err != nil {
			return err
		}
	case address.BookingAddressType:
		if err = s.removeParticipantBooking(ctx, user, *userByPortalCode, addressID); err != nil {
			return err
		}
	case address.UnknownAddressType:
		return ErrUnknownTypeAddress
	default:
		return ErrUnknownTypeAddress
	}

	return nil
}

func (s Service) addParticipantRoom(
	ctx context.Context,
	userCtx, user model.User,
	addressID uuid.UUID,
	role string,
) error {
	roo, err := s.checkAwsRoom(ctx, addressID, userCtx)
	if err != nil {
		return err
	}

	if participantExists(roo.Participants, &user) {
		return room.ErrParticipantExists 
	}

	return s.store.Room().UpdateParticipant(ctx, roo.ID, getParticipantsWithUUID(roo.Participants, &user, role))
}

func (s Service) updateParticipantRoom(
	ctx context.Context,
	userCtx, userBase, userChange model.User,
	addressID uuid.UUID,
	role string,
) error {
	roo, err := s.checkAwsRoom(ctx, addressID, userCtx)
	if err != nil {
		return err
	}

	if userBase.ID != userChange.ID {
		if participantExists(roo.Participants, &userChange) {
			return room.ErrParticipantExists
		}
	}

	roo.Participants = updateParticipant(roo.Participants, userBase, userChange, role)

	return s.store.Room().UpdateParticipant(ctx, roo.ID, getParticipantsWithUUID(roo.Participants, &userChange, role))
}

func (s Service) removeParticipantRoom(
	ctx context.Context,
	userCtx, user model.User,
	addressID uuid.UUID,
) error {
	roo, err := s.checkAwsRoom(ctx, addressID, userCtx)
	if err != nil {
		return err
	}

	return s.
		store.
		Room().
		UpdateParticipant(
			ctx,
			roo.ID,
			getParticipantsWithUUID(removeParticipant(roo.Participants, user), nil, ""),
		)
}

func (s Service) addParticipantBooking(
	ctx context.Context,
	userCtx, user model.User,
	addressID uuid.UUID,
	roleDescription string,
) error {
	book, err := s.checkAwsBooking(ctx, addressID, userCtx)
	if err != nil {
		return err
	}

	return s.store.Booking().CreateBookingUser(ctx, model.BookingUser{
		BookingID:       book.ID,
		UserID:          user.ID,
		Role:            model.ParticipantBookingUserRole,
		RoleDescription: roleDescription,
	})
}

func (s Service) updateParticipantBooking(
	ctx context.Context,
	userCtx, userBase, userChange model.User,
	addressID uuid.UUID,
	roleDescription string,
) error {
	book, err := s.checkAwsBooking(ctx, addressID, userCtx)
	if err != nil {
		return err
	}

	return s.store.Booking().UpdateBookingUser(ctx,
		model.BookingUser{
			BookingID:       book.ID,
			UserID:          userBase.ID,
			Role:            model.ParticipantBookingUserRole,
			RoleDescription: roleDescription,
		},
		model.BookingUser{
			BookingID:       book.ID,
			UserID:          userChange.ID,
			Role:            model.ParticipantBookingUserRole,
			RoleDescription: roleDescription,
		})
}

func (s Service) removeParticipantBooking(
	ctx context.Context,
	userCtx, user model.User,
	addressID uuid.UUID,
) error {
	book, err := s.checkAwsBooking(ctx, addressID, userCtx)
	if err != nil {
		return err
	}

	return s.store.Booking().DeleteBookingUser(ctx, model.BookingUser{
		BookingID: book.ID,
		UserID:    user.ID,
		Role:      model.ParticipantBookingUserRole,
	})
}

func participantExists(participants []model.Participant, newParticipant *model.User) bool {
	for _, user := range participants {
		if user.User.ID == newParticipant.ID {
			return true
		}
	}

	return false
}

func getParticipantsWithUUID(participants []model.Participant, newParticipant *model.User, newRole string) map[uuid.UUID]string {
	participantsUUID := make(map[uuid.UUID]string, len(participants))

	for _, user := range participants {
		participantsUUID[user.User.ID] = user.Role
	}

	if newParticipant != nil {
		participantsUUID[newParticipant.ID] = newRole
	}

	return participantsUUID
}

func updateParticipant(participants []model.Participant, userBase, userChange model.User, newRole string) []model.Participant {
	participantsUpd := make([]model.Participant, 0, len(participants))

	for _, participant := range participants {
		if toChangeParticipant(participant.User, userBase, userChange) {
			participantsUpd = removeParticipant(participants, userBase)

			participantsUpd = append(participantsUpd, model.Participant{
				User: userChange,
				Role: newRole,
			})

			break
		}

		if toChangeRole(participant.User, userBase, userChange) {
			participantsUpd = participants

			participantsUpd = append(participantsUpd, model.Participant{
				User: userChange,
				Role: newRole,
			})

			break
		}
	}

	if len(participantsUpd) == 0 {
		participantsUpd = participants
	}

	return participantsUpd
}

func toChangeParticipant(participant, userBase, userChange model.User) bool {
	return participant.ID == userBase.ID && userBase.ID != userChange.ID
}

func toChangeRole(participant, userBase, userChange model.User) bool {
	return participant.ID == userBase.ID && userBase.ID == userChange.ID
}

func removeParticipant(participants []model.Participant, user model.User) []model.Participant {
	newParticipant := make([]model.Participant, 0, len(participants))

	for _, participant := range participants {
		if participant.User.ID == user.ID {
			continue
		}

		newParticipant = append(newParticipant, participant)
	}

	return newParticipant
}

func (s Service) checkAwsRoom(ctx context.Context, roomID uuid.UUID, user model.User) (*model.Room, error) {
	roo, err := s.store.Room().FindByID(ctx, roomID)
	if err != nil {
		return nil, room.ErrRoomNotFound
	}

	if !acl.CanAddParticipants(user, roo) {
		return nil, ErrPermissionDenied
	}

	return roo	, nil
}

func (s Service) checkAwsBooking(ctx context.Context, bookingID uuid.UUID, user model.User) (*model.Booking, error) {
	book, err := s.store.Booking().FindByID(ctx, bookingID)
	if err != nil {
		return nil, booking.ErrBookingNotFound
	}

	if !acl.CanAddParticipants(user, book) {
		return nil, ErrPermissionDenied
	}

	return book, nil
}
