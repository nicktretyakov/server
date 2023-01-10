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

func (s Service) AddAttachment(
	ctx context.Context,
	user model.User,
	attachment model.Attachment,
	typeAddress address.TypeAddress,
) (*model.Attachment, error) {
	var (
		book *model.Booking
		roo *model.Room
		err  error
	)

	switch typeAddress {
	case address.BookingAddressType:
		book, err = s.store.Booking().FindByID(ctx, attachment.AddressID)
		if err != nil {
			return nil, booking.ErrBookingNotFound
		}

		if !acl.CanAddAttachment(user, book) {
			return nil, ErrPermissionDenied
		}
	case address.RoomAddressType:
		roo, err = s.store.Room().FindByID(ctx, attachment.AddressID)
		if err != nil {
			return nil, room.ErrRoomNotFound
		}

		if !acl.CanAddAttachment(user, roo) {
			return nil, ErrPermissionDenied
		}
	case address.UnknownAddressType:
		return nil, ErrUnknownTypeAddress
	default:
		return nil, ErrUnknownTypeAddress
	}

	attachment.AuthorID = user.ID

	return s.store.Address().CreateAttachment(ctx, attachment)
}

func (s Service) RemoveAttachment(
	ctx context.Context,
	user model.User,
	attachmentID uuid.UUID,
	typeAddress address.TypeAddress,
) error {
	var (
		book *model.Booking
		roo *model.Room
		err  error
	)

	attachment, err := s.store.Address().FindAttachmentByID(ctx, attachmentID)
	if err != nil {
		return ErrAttachmentNotFound
	}

	switch typeAddress {
	case address.BookingAddressType:
		book, err = s.store.Booking().FindByID(ctx, attachment.AddressID)
		if err != nil {
			return booking.ErrBookingNotFound
		}

		if !acl.CanRemoveAttachment(user, book) {
			return ErrPermissionDenied
		}
	case address.RoomAddressType:
		roo, err = s.store.Room().FindByID(ctx, attachment.AddressID)
		if err != nil {
			return room.ErrRoomNotFound
		}

		if !acl.CanRemoveAttachment(user, roo) {
			return ErrPermissionDenied
		}
	case address.UnknownAddressType:
		return ErrUnknownTypeAddress
	default:
		return ErrUnknownTypeAddress
	}

	return s.store.Address().DeleteAttachment(ctx, *attachment)
}

func (s Service) RenameAttachment(
	ctx context.Context, user model.User, attachmentID uuid.UUID, newFileName string, typeAddress address.TypeAddress,
) (*model.Attachment, error) {
	var (
		book *model.Booking
		roo *model.Room
		err  error
	)

	attachment, err := s.store.Address().FindAttachmentByID(ctx, attachmentID)
	if err != nil {
		return nil, ErrAttachmentNotFound
	}

	switch typeAddress {
	case address.BookingAddressType:
		book, err = s.store.Booking().FindByID(ctx, attachment.AddressID)
		if err != nil {
			return nil, booking.ErrBookingNotFound
		}

		if !acl.CanRenameAttachment(user, book) {
			return nil, ErrPermissionDenied
		}
	case address.RoomAddressType:
		roo, err = s.store.Room().FindByID(ctx, attachment.AddressID)
		if err != nil {
			return nil, room.ErrRoomNotFound
		}

		if !acl.CanRenameAttachment(user, roo) {
			return nil, ErrPermissionDenied
		}
	case address.UnknownAddressType:
		return nil, ErrUnknownTypeAddress
	default:
		return nil, ErrUnknownTypeAddress
	}

	attachment.Filename = newFileName

	return s.store.Address().RenameAttachment(ctx, *attachment)
}
