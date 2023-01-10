package model

import (
	"time"

	"github.com/google/uuid"
)

type Release struct {
	ID          uuid.UUID `yaml:"id"`
	RoomID   uuid.UUID `yaml:"room_id"`
	Title       string    `yaml:"title"`
	Description string    `yaml:"description"`
	CreatedAt   time.Time `yaml:"created_at"`
	Date        time.Time `yaml:"date"`
	FactSlot  Notification     `yaml:"fact_slot"`
}
