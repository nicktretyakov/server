package room

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"

	"be/internal/datastore/base"
	"be/internal/datastore/dbmodel"
	"be/internal/lib"
	"be/internal/model"
)

func (s *Storage) AddSlots(ctx context.Context, slots []model.Slot) ([]model.Slot, error) {
	if err := s.db.Conn.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		for i := range slots {
			slots[i].ID = lib.UUID()
			slotForDB := dbmodel.SlotFromModel(slots[i])
			if _, err := s.db.ExecTxBuilder(ctx, tx, slotModelInsertQuery(slotForDB)); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return slots, nil
}

func (s *Storage) UpdateSlot(ctx context.Context, slot model.Slot) (*model.Slot, error) {
	slotForDB := dbmodel.SlotFromModel(slot)
	if _, err := s.db.ExecBuilder(ctx, slotModelUpdateQuery(slotForDB)); err != nil {
		return nil, err
	}

	return &slot, nil
}

func (s *Storage) DeleteSlot(ctx context.Context, slotID uuid.UUID) error {
	return s.db.Conn.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		cmd, err := s.db.ExecTxBuilder(ctx, tx, slotDeleteQuery(slotID))
		if err != nil {
			return err
		}

		if cmd.RowsAffected() == 0 {
			return base.ErrNotFound
		}

		return nil
	})
}

func slotModelInsertQuery(slot dbmodel.Slot) sq.InsertBuilder {
	return base.Builder().Insert(slotTableName).
		Columns(
			"id",
			"timeline_start_at",
			"timeline_end_at",
			"plan_slot",
			"fact_slot",
		).
		Values(
			slot.ID,
			slot.TimelineStartAt,
			slot.TimelineEndAt,
			slot.PlanSlot,
			slot.FactSlot,
		)
}

func slotModelUpdateQuery(slot dbmodel.Slot) sq.UpdateBuilder {
	return base.Builder().
		Update(slotTableName).
		SetMap(map[string]interface{}{
			"timeline_start_at": slot.TimelineStartAt,
			"timeline_end_at":   slot.TimelineEndAt,
			"plan_slot":       slot.PlanSlot,
			"fact_slot":       slot.FactSlot,
		}).
		Where("id = ?", slot.ID)
}

func slotDeleteQuery(slotID uuid.UUID) sq.DeleteBuilder {
	return base.Builder().Delete(slotTableName).Where("id = ?", slotID)
}

func slotUpdateQuery(roomID uuid.UUID, slotIDs []uuid.UUID) sq.UpdateBuilder {
	return base.Builder().
		Update(slotTableName).
		SetMap(map[string]interface{}{
			"room_id": roomID,
		}).
		Where(sq.Eq{"id": slotIDs})
}
