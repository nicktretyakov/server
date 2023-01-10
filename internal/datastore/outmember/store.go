package outmember

import (
	"be/internal/datastore/base"
	addressStore "be/internal/datastore/address"
)

const (
	BookingOutmemberTableName = "booking_outmember"
	RoomOutmemberTableName = "room_outmember"
)

type storage struct {
	db *base.DB
}

func New(db *base.DB) *storage {
	return &storage{db: db}
}

func getNameTableQuery(typeAddress addressStore.TypeAddress) string {
	switch typeAddress {
	case addressStore.BookingAddressType:
		return BookingOutmemberTableName
	case addressStore.RoomAddressType:
		return RoomOutmemberTableName
	case addressStore.UnknownAddressType:
		return ""
	default:
		return ""
	}
}
