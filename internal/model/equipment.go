package model

import (
	"time"

	"github.com/google/uuid"
)

type Equipment struct {
	ID          uuid.UUID `yaml:"id"`
	RoomID   uuid.UUID `yaml:"room_id"`
	Title       string    `yaml:"title"`
	Timeline    Timeline  `yaml:"timeline"`
	Description string    `yaml:"description"`
	PlanValue   float32   `yaml:"plan_value"`
	FactValue   float32   `yaml:"fact_value"`
	CreatedAt   time.Time `yaml:"created_at"`
}
