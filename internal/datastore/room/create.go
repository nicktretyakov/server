package room

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"

	"be/internal/datastore/base"
	"be/internal/datastore/dbmodel"
	"be/internal/lib"
	"be/internal/model"
)

var ErrEmptySupervisor = errors.New("empty supervisor")

//nolint:gocognit
func (s *Storage) Create(ctx context.Context, roomToStore model.Room, equipmentIDs, slotIDs []uuid.UUID) (*model.Room, error) {
	if roomToStore.ID != uuid.Nil {
		return nil, errors.New("already exists")
	}

	roomToStore.ID = lib.UUID()
	roomToStore.CreatedAt = s.db.Now()
	roomToStore.UpdatedAt = roomToStore.CreatedAt

	return &roomToStore, s.db.Conn.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		dbRoom := dbmodel.RoomFromModel(roomToStore)
		if err := s.db.GetTx(ctx, tx, roomInsertQuery(dbRoom), &roomToStore.Number); err != nil {
			return err
		}

		if len(equipmentIDs) > 0 {
			if _, err := s.db.ExecTxBuilder(ctx, tx, equipmentUpdateQuery(roomToStore.ID, equipmentIDs)); err != nil {
				return err
			}
		}

		if len(slotIDs) > 0 {
			if _, err := s.db.ExecTxBuilder(ctx, tx, slotUpdateQuery(roomToStore.ID, slotIDs)); err != nil {
				return err
			}
		}

		return nil
	})
}

func roomInsertQuery(p dbmodel.Room) sq.InsertBuilder {
	return base.Builder().Insert(roomTableName).
		Columns(
			"id",
			"created_at",
			"updated_at",
			"author_id",
			"title",
			"description",
			"target_audience",
			"status",
			"links",
			"employee_id",
			"owner_id",
			"booking_ids",
			"creation_date",
			"space",
	        "security_email",
	        "visible",
		).
		Values(
			p.ID,             // id
			p.CreatedAt,      // created_at
			p.UpdatedAt,      // updated_at
			p.Author.ID,      // author_id
			p.Title,          // title
			p.Description,    // description
			p.TargetAudience, // target_audience
			p.Status,         // status
			p.Links,          // links
			p.Employee.ID,    // employee_id
			p.Owner.ID,       // owner_id
			p.Bookings,       // booking_ids
			p.CreationDate,   // creation_date
		).
		Suffix("RETURNING number")
}
