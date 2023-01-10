package address

import (
	"time"

	"be/internal/datastore/base"
	"be/internal/filestorage"
)

const (
	attachmentsTableName        = "attachments"
	linkTimeLife                = time.Hour
	emailNoteTableName  = "email_note"
	lifeNoteTableName   = "life_note"
	systemNoteTableName = "system_note"
)

type TypeAddress int

const (
	UnknownAddressType TypeAddress = iota
	BookingAddressType
	RoomAddressType
)

type Storage struct {
	db           *base.DB
	linkTimeLife time.Duration
	fileStorage  filestorage.IFileStorage
}

func New(db *base.DB, fileStorage filestorage.IFileStorage) *Storage {
	return &Storage{db: db, fileStorage: fileStorage, linkTimeLife: linkTimeLife}
}
