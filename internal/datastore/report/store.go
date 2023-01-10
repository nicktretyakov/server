package report

import (
	"be/internal/datastore/base"
)

const (
	reportBookingTableName = "booking_reports"
	reportRoomTableName = "room_reports"
	reportTableNameAlias   = "booking_reports reports"
)

type Storage struct {
	db *base.DB
}

func New(db *base.DB) *Storage {
	return &Storage{db: db}
}
