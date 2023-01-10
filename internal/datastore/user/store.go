package user

import "be/internal/datastore/base"

const (
	userTableName                 = "users"
	noteSettingsTableName = "note_settings"
)

type Storage struct {
	db *base.DB
}

func New(db *base.DB) *Storage {
	return &Storage{db: db}
}
