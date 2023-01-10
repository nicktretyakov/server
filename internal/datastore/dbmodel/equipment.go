package dbmodel

import (
	"time"

	"github.com/google/uuid"

	"be/internal/model"
)

type (
	Equipment struct {
		ID              uuid.UUID `yaml:"id"`
		RoomID       uuid.UUID `yaml:"room_id"`
		Title           string    `yaml:"title"`
		TimelineStartAt time.Time `yaml:"timeline_start_at"`
		TimelineEndAt   time.Time `yaml:"timeline_end_at"`
		Description     string    `yaml:"description"`
		PlanValue       float32   `yaml:"plan_value"`
		FactValue       float32   `yaml:"fact_value"`
		CreatedAt       time.Time `yaml:"created_at"`
	}

	EquipmentList []Equipment
)

func EquipmentFromModel(m model.Equipment) Equipment {
	return Equipment{
		ID:              m.ID,
		RoomID:       m.RoomID,
		TimelineStartAt: m.Timeline.StartAt,
		TimelineEndAt:   m.Timeline.EndAt,
		Title:           m.Title,
		Description:     m.Description,
		PlanValue:       m.PlanValue,
		FactValue:       m.FactValue,
		CreatedAt:       m.CreatedAt,
	}
}

func (m Equipment) ToModel() model.Equipment {
	slot := model.Equipment{
		ID:        m.ID,
		RoomID: m.RoomID,
		Timeline: model.Timeline{
			StartAt: m.TimelineStartAt,
			EndAt:   m.TimelineEndAt,
		},
		Title:       m.Title,
		Description: m.Description,
		PlanValue:   m.PlanValue,
		FactValue:   m.FactValue,
		CreatedAt:   m.CreatedAt,
	}

	return slot
}

func ToEquipmentList(equipments []model.Equipment) EquipmentList {
	equipmentList := make(EquipmentList, 0, len(equipments))
	for _, equipment := range equipments {
		equipmentList = append(equipmentList, EquipmentFromModel(equipment))
	}

	return equipmentList
}

func (m EquipmentList) Equipments() []model.Equipment {
	equipments := make([]model.Equipment, 0, len(m))
	for _, equipment := range m {
		equipments = append(equipments, equipment.ToModel())
	}

	return equipments
}

func (m EquipmentList) EquipmentsID() []uuid.UUID {
	equipmentUUIDs := make([]uuid.UUID, 0, len(m))

	for _, equipment := range m {
		equipmentUUIDs = append(equipmentUUIDs, equipment.ID)
	}

	return equipmentUUIDs
}
