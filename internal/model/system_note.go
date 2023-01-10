package model

import (
	"time"

	"github.com/google/uuid"
)

type SystemStatus string

const (
	NotRead SystemStatus = "not_read"
	Read    SystemStatus = "read"
)

type SystemNote struct {
	ID             uuid.UUID          `db:"id" yaml:"id"`
	Event          NoteEvent  `db:"event" yaml:"event"`
	Status         SystemStatus       `db:"status" yaml:"status"`
	Object         NoteObject `db:"object" yaml:"object"`
	AddressID *uuid.UUID         `db:"address_id" yaml:"address_id"`
	ActorID        uuid.UUID          `db:"actor_id" yaml:"actor_id"`
	RecipientID    uuid.UUID          `db:"recipient_id" yaml:"recipient_id"`
	Header         string             `db:"header" yaml:"header"`
	Body           string             `db:"body" yaml:"body"`
	CreatedAt      time.Time          `db:"created_at" yaml:"created_at"`
	ReadAt         time.Time          `db:"read_at" yaml:"read_at"`
}
