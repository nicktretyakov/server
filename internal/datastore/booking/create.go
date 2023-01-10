package booking

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

func (s *Storage) Create(ctx context.Context, bookingToStore model.Booking) (*model.Booking, error) {
	if bookingToStore.ID != uuid.Nil {
		return nil, errors.New("already exists")
	}

	if bookingToStore.Supervisor == nil {
		return nil, ErrEmptySupervisor
	}

	if bookingToStore.Author == nil {
		return nil, ErrEmptySupervisor
	}

	bookingToStore.ID = lib.UUID()
	bookingToStore.CreatedAt = s.db.Now()
	bookingToStore.UpdatedAt = bookingToStore.CreatedAt

	return &bookingToStore, s.db.Conn.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		insertQuery := bookingInsertQuery(dbmodel.BookingFromModel(bookingToStore))
		if err := s.db.GetTx(ctx, tx, insertQuery, &bookingToStore.Number); err != nil {
			return err
		}

		if _, err := s.db.ExecTxBuilder(ctx, tx, bookingUserInsertQuery(model.BookingUser{
			BookingID: bookingToStore.ID,
			UserID:    bookingToStore.Supervisor.ID,
			Role:      model.SupervisorBookingUserRole,
		})); err != nil {
			return err
		}

		_, err := s.db.ExecTxBuilder(ctx, tx, bookingUserInsertQuery(model.BookingUser{
			BookingID: bookingToStore.ID,
			UserID:    bookingToStore.Author.ID,
			Role:      model.AuthorBookingUserRole,
		}))

		return err
	})
}

func bookingInsertQuery(p dbmodel.Booking) sq.InsertBuilder {
	return base.Builder().Insert(bookingTableName).
		Columns(
			"id",
			"created_at",
			"updated_at",
			"title",
			"city",
			"timeline_start_at",
			"timeline_end_at",
			"description",
			"status",
			"slot",
			"goal",
			"type",
			"links",
		).
		Values(
			p.ID,              // id
			p.CreatedAt,       // created_at
			p.UpdatedAt,       // updated_at
			p.Title,           // title
			p.City,            // city
			p.TimelineStartAt, // timeline
			p.TimelineEndAt,   // timeline
			p.Description,     // description
			p.Status,          // status
			p.Slot,          // slot
			p.Goal,            // goal
			p.Type,            // type
			p.Links,           // links
		).
		Suffix("RETURNING number")
}
