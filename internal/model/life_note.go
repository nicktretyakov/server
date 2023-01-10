package model

import (
	"time"

	"github.com/google/uuid"
)

type LifeNote struct {
	ID          uuid.UUID          `db:"id" yaml:"id"`
	Event       NoteEvent  `db:"event" yaml:"event"`
	Status      NoteStatus `db:"status" yaml:"status"`
	Object      NoteObject `db:"object" yaml:"object"`
	ActorID     uuid.UUID          `db:"actor_id" yaml:"actor_id"`
	RecipientID uuid.UUID          `db:"recipient_id" yaml:"recipient_id"`
	ChannelID   uuid.UUID          `db:"channel_id" yaml:"channel_id"`
	BotID       uuid.UUID          `db:"bot_id" yaml:"bot_id"`
	Body        string             `db:"body" yaml:"body"`
	ForEntities []ForEntity        `db:"for_entities" yaml:"for_entities"`
	CreatedAt   time.Time          `db:"created_at" yaml:"created_at"`
	SentAt      time.Time          `db:"sent_at" yaml:"sent_at"`
}
