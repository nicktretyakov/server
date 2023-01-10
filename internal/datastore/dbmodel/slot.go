package dbmodel

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"

	"be/internal/model"
)

type (
	Slot struct {
		ID              uuid.UUID      `db:"id"`
		RoomID       uuid.UUID      `db:"room_id"`
		TimelineStartAt time.Time      `db:"timeline_start_at"`
		TimelineEndAt   time.Time      `db:"timeline_end_at"`
		PlanSlot      pgtype.Numeric `db:"plan_slot"`
		FactSlot      pgtype.Numeric `db:"fact_slot"`
		CreatedAt       time.Time      `db:"created_at"`
	}

	SlotList []Slot
)

func SlotFromModel(b model.Slot) Slot {
	planSlot := pgtype.Numeric{}
	_ = planSlot.Set(b.PlanSlot.String())

	factSlot := pgtype.Numeric{}
	_ = factSlot.Set(b.FactSlot.String())

	return Slot{
		ID:              b.ID,
		RoomID:       b.RoomID,
		TimelineStartAt: b.Timeline.StartAt,
		TimelineEndAt:   b.Timeline.EndAt,
		PlanSlot:      planSlot,
		FactSlot:      factSlot,
		CreatedAt:       b.CreatedAt,
	}
}

func (p Slot) ToModel() model.Slot {
	slot := model.Slot{
		ID:        p.ID,
		RoomID: p.RoomID,
		Timeline: model.Timeline{
			StartAt: p.TimelineStartAt,
			EndAt:   p.TimelineEndAt,
		},
		PlanSlot: ToNotification(p.PlanSlot),
		FactSlot: ToNotification(p.FactSlot),
		CreatedAt:  p.CreatedAt,
	}

	return slot
}

func ToSlotList(slots []model.Slot) SlotList {
	slotList := make(SlotList, 0, len(slots))
	for _, slot := range slots {
		slotList = append(slotList, SlotFromModel(slot))
	}

	return slotList
}

func (b SlotList) Slots() []model.Slot {
	slots := make([]model.Slot, 0, len(b))
	for _, slot := range b {
		slots = append(slots, slot.ToModel())
	}

	return slots
}

func (b SlotList) SlotsID() []uuid.UUID {
	slotUUIDs := make([]uuid.UUID, 0, len(b))

	for _, slot := range b {
		slotUUIDs = append(slotUUIDs, slot.ID)
	}

	return slotUUIDs
}
