package room

import (
	"time"

	"be/internal/datastore/base"
	"be/internal/filestorage"
)

const (
	roomTableName      = "rooms"
	equipmentTableName       = "room_equipment"
	slotTableName       = "room_slot"
	roomTableNameAlias = "rooms as p"
	attachmentsTableName  = "attachments"
	releaseTableName      = "room_release"
	reportTableName       = "room_reports"

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
