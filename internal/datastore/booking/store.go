package booking

import (
	"time"

	"be/internal/datastore/base"
	"be/internal/filestorage"
)

const (
	bookingUserRoleTableName = "booking_users"
	bookingReportsTableName  = "booking_reports"
	bookingTableName         = "bookings"
	bookingTableNameAlias    = "bookings as p"
	attachmentsTableName     = "attachments"

	linkTimeLife = time.Hour
)

type Storage struct {
	db           *base.DB
	linkTimeLife time.Duration
	fileStorage  filestorage.IFileStorage
}

func New(db *base.DB, fileStorage filestorage.IFileStorage) *Storage {
	return &Storage{db: db, fileStorage: fileStorage, linkTimeLife: linkTimeLife}
}
