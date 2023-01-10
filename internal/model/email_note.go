package model

import (
	"time"

	"github.com/google/uuid"
)

type EmailNote struct {
	ID             uuid.UUID          `db:"id" yaml:"id"`
	Event          NoteEvent  `db:"event" yaml:"event"`
	Status         NoteStatus `db:"status" yaml:"status"`
	ActorID        uuid.UUID          `db:"actor_id" yaml:"actor_id"`
	Object         NoteObject `db:"object" yaml:"object"`
	RecipientID    uuid.UUID          `db:"recipient_id" yaml:"recipient_id"`
	RecipientEmail string             `db:"recipient_email" yaml:"recipient_email"`
	SenderEmail    string             `db:"sender_email" yaml:"sender_email"`
	Subject        string             `db:"subject" yaml:"subject"`
	Body           string             `db:"body" yaml:"body"`
	CreatedAt      time.Time          `db:"created_at" yaml:"created_at"`
	SentAt         time.Time          `db:"sent_at" yaml:"sent_at"`
}

const (
	LinkPattern string = "[[link]]"
)

type ForEntity struct {
	Pattern string
	Value   string
	Link    string
}
