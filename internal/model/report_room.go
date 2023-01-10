package model

import (
	"time"

	"github.com/google/uuid"
)

type ReportRoom struct {
	ID        uuid.UUID `yaml:"id"`
	CreatedAt time.Time `yaml:"created_at"`
	UpdatedAt time.Time `yaml:"updated_at"`
	RoomID uuid.UUID `yaml:"room_id"`
	Comment   *string   `yaml:"comment"`
	Timeline  Timeline  `yaml:"timeline"`
	Equipments   []Equipment  `yaml:"equipments"`
	Slots   []Slot  `yaml:"slots"`
	Releases  []Release `yaml:"releases"`
}
