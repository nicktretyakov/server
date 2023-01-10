package model

import (
	"time"

	"github.com/google/uuid"
)

type Slot struct {
	ID         uuid.UUID `yaml:"id"`
	RoomID  uuid.UUID `yaml:"room_id"`
	Timeline   Timeline  `yaml:"timeline"`
	PlanSlot Notification     `yaml:"plan_slot"`
	FactSlot Notification     `yaml:"fact_slot"`
	CreatedAt  time.Time `yaml:"created_at"`
}
