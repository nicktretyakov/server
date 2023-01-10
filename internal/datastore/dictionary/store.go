package dictionary

import (
	"be/internal/datastore/base"
)

const (
	DepartmentsTableName      = "dictionary_departments"
	DepartmentsTableNameAlias = "dictionary_departments as dep"
)

type storage struct {
	db *base.DB
}

func New(db *base.DB) *storage {
	return &storage{db: db}
}
