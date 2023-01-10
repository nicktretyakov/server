package finalreports

import (
	"be/internal/datastore/base"
)

const (
	finalReportsTableName      = "booking_final_reports"
	finalReportsTableNameAlias = "booking_final_reports final_reports"
)

type storage struct {
	db *base.DB
}

func New(db *base.DB) *storage {
	return &storage{db: db}
}
