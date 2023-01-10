package address

import (
	"context"
	"github.com/google/uuid"
	"be/internal/acl"
	"be/internal/booking"
	"be/internal/datastore/address"
	"be/internal/lib"
	"be/internal/model"
	"be/internal/room"
)

func (s Service) AddLink(
	ctx context.Context,
	user model.User,
	link *model.Link,
	addressID uuid.UUID,
	typeAddress address.TypeAddress,
) error {
	switch typeAddress {
	case address.BookingAddressType:
		return s.addLinkBooking(ctx, user, link, addressID)
	case address.RoomAddressType:
		return s.addLinkRoom(ctx, user, link, addressID)
	case address.UnknownAddressType:
		return ErrUnknownTypeAddress
	default:
		return ErrUnknownTypeAddress
	}
}

func (s Service) UpdateLink(
	ctx context.Context,
	user model.User,
	link *model.Link,
	addressID uuid.UUID,
	typeAddress address.TypeAddress,
) error {
	switch typeAddress {
	case address.BookingAddressType:
		return s.updateLinkBooking(ctx, user, link, addressID)
	case address.RoomAddressType:
		return s.updateLinkRoom(ctx, user, link, addressID)
	case address.UnknownAddressType:
		return ErrUnknownTypeAddress
	default:
		return ErrUnknownTypeAddress
	}
}

func (s Service) RemoveLink(
	ctx context.Context,
	user model.User,
	link uuid.UUID,
	addressID uuid.UUID,
	typeAddress address.TypeAddress,
) error {
	switch typeAddress {
	case address.BookingAddressType:
		return s.removeLinkBooking(ctx, user, link, addressID)
	case address.RoomAddressType:
		return s.removeLinkRoom(ctx, user, link, addressID)
	case address.UnknownAddressType:
		return ErrUnknownTypeAddress
	default:
		return ErrUnknownTypeAddress
	}
}

func (s Service) addLinkBooking(ctx context.Context, user model.User, link *model.Link, addressID uuid.UUID) error {
	book, err := s.store.Booking().FindByID(ctx, addressID)
	if err != nil {
		return booking.ErrBookingNotFound
	}

	if !acl.CanAddLink(user, book) {
		return ErrPermissionDenied
	}

	if link.Id == uuid.Nil {
		link.Id = lib.UUID()
	}

	book.Links = append(book.Links, *link)

	return s.store.Booking().UpdateLinks(ctx, book.ID, book.Links)
}

func (s Service) addLinkRoom(ctx context.Context, user model.User, link *model.Link, addressID uuid.UUID) error {
	roo, err := s.store.Room().FindByID(ctx, addressID)
	if err != nil {
		return room.ErrRoomNotFound
	}

	if !acl.CanAddLink(user, roo) {
		return ErrPermissionDenied
	}

	if link.Id == uuid.Nil {
		link.Id = lib.UUID()
	}

	roo.Links = append(roo.Links, *link)

	return s.store.Room().UpdateLinks(ctx, roo.ID, roo.Links)
}

func (s Service) updateLinkBooking(ctx context.Context, user model.User, link *model.Link, addressID uuid.UUID) error {
	book, err := s.store.Booking().FindByID(ctx, addressID)
	if err != nil {
		return booking.ErrBookingNotFound
	}

	if !acl.CanUpdateLink(user, book) {
		return ErrPermissionDenied
	}

	updateLink(link, &book.Links)

	return s.store.Booking().UpdateLinks(ctx, book.ID, book.Links)
}

func (s Service) updateLinkRoom(ctx context.Context, user model.User, link *model.Link, addressID uuid.UUID) error {
	roo, err := s.store.Room().FindByID(ctx, addressID)
	if err != nil {
		return room.ErrRoomNotFound
	}

	if !acl.CanUpdateLink(user, roo) {
		return ErrPermissionDenied
	}

	updateLink(link, &roo.Links)

	return s.store.Room().UpdateLinks(ctx, roo.ID, roo.Links)
}

func updateLink(linkForUpdate *model.Link, links *[]model.Link) {
	for i := range *links {
		if (*links)[i].Id == linkForUpdate.Id {
			(*links)[i] = *linkForUpdate

			break
		}
	}
}

func (s Service) removeLinkBooking(ctx context.Context, user model.User, link uuid.UUID, addressID uuid.UUID) error {
	book, err := s.store.Booking().FindByID(ctx, addressID)
	if err != nil {
		return booking.ErrBookingNotFound
	}

	if !acl.CanRemoveLink(user, book) {
		return ErrPermissionDenied
	}

	book.Links = removeLink(link, book.Links)

	return s.store.Booking().UpdateLinks(ctx, book.ID, book.Links)
}

func (s Service) removeLinkRoom(ctx context.Context, user model.User, link uuid.UUID, addressID uuid.UUID) error {
	roo, err := s.store.Room().FindByID(ctx, addressID)
	if err != nil {
		return room.ErrRoomNotFound
	}

	if !acl.CanRemoveLink(user, *roo) {
		return ErrPermissionDenied
	}

	roo.Links = removeLink(link, roo.Links)

	return s.store.Room().UpdateLinks(ctx, roo.ID, roo.Links)
}

func removeLink(linkForRemove uuid.UUID, links []model.Link) []model.Link {
	linksAfterRemove := make([]model.Link, 0, len(links))

	for _, link := range links {
		if linkForRemove == link.Id {
			continue
		}

		linksAfterRemove = append(linksAfterRemove, link)
	}

	return linksAfterRemove
}
