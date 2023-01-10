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

func (s *Storage) AddEquipments(ctx context.Context, equipments []model.Equipment) ([]model.Equipment, error) {
	if err := s.db.Conn.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		for i := range equipments {
			equipments[i].ID = lib.UUID()
			equipmentForDB := dbmodel.EquipmentFromModel(equipments[i])
			if _, err := s.db.ExecTxBuilder(ctx, tx, equipmentInsertQuery(equipmentForDB)); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return equipments, nil
}

func (s *Storage) UpdateEquipment(ctx context.Context, equipment model.Equipment) (*model.Equipment, error) {
	equipmentForDB := dbmodel.EquipmentFromModel(equipment)
	if _, err := s.db.ExecBuilder(ctx, equipmentModelUpdateQuery(equipmentForDB)); err != nil {
		return nil, err
	}

	return &equipment, nil
}

func (s *Storage) DeleteEquipment(ctx context.Context, equipmentID uuid.UUID) error {
	return s.db.Conn.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		cmd, err := s.db.ExecTxBuilder(ctx, tx, equipmentDeleteQuery(equipmentID))
		if err != nil {
			return err
		}

		if cmd.RowsAffected() == 0 {
			return base.ErrNotFound
		}

		return nil
	})
}

func equipmentInsertQuery(equipment dbmodel.Equipment) sq.InsertBuilder {
	return base.Builder().Insert(equipmentTableName).
		Columns(
			"id",
			"title",
			"timeline_start_at",
			"timeline_end_at",
			"description",
			"plan_value",
			"fact_value",
		).
		Values(
			equipment.ID,
			equipment.Title,
			equipment.TimelineStartAt,
			equipment.TimelineEndAt,
			equipment.Description,
			equipment.PlanValue,
			equipment.FactValue,
		)
}

func equipmentModelUpdateQuery(equipment dbmodel.Equipment) sq.UpdateBuilder {
	return base.Builder().
		Update(equipmentTableName).
		SetMap(map[string]interface{}{
			"title":             equipment.Title,
			"timeline_start_at": equipment.TimelineStartAt,
			"timeline_end_at":   equipment.TimelineEndAt,
			"description":       equipment.Description,
			"plan_value":        equipment.PlanValue,
			"fact_value":        equipment.FactValue,
		}).
		Where("id = ?", equipment.ID)
}

func equipmentDeleteQuery(equipmentID uuid.UUID) sq.DeleteBuilder {
	return base.Builder().Delete(equipmentTableName).Where("id = ?", equipmentID)
}

func equipmentUpdateQuery(roomID uuid.UUID, equipmentIDs []uuid.UUID) sq.UpdateBuilder {
	return base.Builder().
		Update(equipmentTableName).
		SetMap(map[string]interface{}{
			"room_id": roomID,
		}).
		Where(sq.Eq{"id": equipmentIDs})
}
