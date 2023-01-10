package model

import (
	"time"

	"github.com/google/uuid"
)

type AggregateStatus int

const (
	ActiveAggregateStatus  AggregateStatus = 1
	InitialAggregateStatus AggregateStatus = 2
)

type Stage struct {
	ID        uuid.UUID
	Status    AggregateStatus
	CreatedAt time.Time

	BookingID uuid.UUID
	Title     string
	Timeline  *Timeline
	Issues    []Issue
}

type Issue struct {
	ID           uuid.UUID
	Status       AggregateStatus
	Stage        Stage
	CreatedAt    time.Time
	Title        string
	Description  string
	Timeline     *Timeline
	Participants []User
	Attachments  []Attachment
}
