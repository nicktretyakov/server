package session

import (
	"be/internal/datastore/base"
)

const sessionTableName = "sessions"

type storage struct {
	db *base.DB
}

func New(db *base.DB) *storage {
	return &storage{db: db}
}
