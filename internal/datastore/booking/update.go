package booking

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"

	"be/internal/datastore/base"
	"be/internal/datastore/dbmodel"
	"be/internal/model"
)

func (s *Storage) UpdateStatus(ctx context.Context, bookingID uuid.UUID, status model.Status) error {
	book := dbmodel.Booking{Status: status, ID: bookingID}
	book.UpdatedAt = s.db.Now()

	sql := bookingUpdateStatusQuery(book)
	if cmd, err := s.db.ExecBuilder(ctx, sql); err != nil || cmd.RowsAffected() == 0 {
		if err != nil {
			return err
		}

		return base.ErrNotFound
	}

	return nil
}

func (s *Storage) Update(ctx context.Context, book model.Booking) (*model.Booking, error) {
	bookingToStore := dbmodel.BookingFromModel(book)
	bookingToStore.UpdatedAt = s.db.Now()

	if err := s.db.Conn.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		if _, err := s.db.ExecTxBuilder(ctx, tx, bookingUpdateQuery(bookingToStore)); err != nil {
			return err
		}

		// delete old departments
		if _, err := s.db.ExecTxBuilder(ctx, tx, bookingDepartmentsDeleteQuery(bookingToStore)); err != nil {
			return err
		}

		if len(bookingToStore.Departments) > 0 {
			// insert new departments
			if _, err := s.db.ExecTxBuilder(ctx, tx, bookingDepartmentsInsertQuery(bookingToStore)); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return bookingToStore.ToModelPtr(), nil
}

func (s *Storage) UpdateLinks(ctx context.Context, bookingID uuid.UUID, links []model.Link) error {
	book := dbmodel.Booking{Links: links, ID: bookingID}
	book.UpdatedAt = s.db.Now()

	if cmd, err := s.db.ExecBuilder(ctx, bookingUpdateLinksQuery(book)); err != nil || cmd.RowsAffected() == 0 {
		if err != nil {
			return err
		}

		return base.ErrNotFound
	}

	return nil
}

func (s *Storage) UpdateState(ctx context.Context, bookingID uuid.UUID, state model.State) error {
	book := dbmodel.Booking{State: state, ID: bookingID}
	book.UpdatedAt = s.db.Now()

	if cmd, err := s.db.ExecBuilder(ctx, bookingUpdateStateQuery(book)); err != nil || cmd.RowsAffected() == 0 {
		if err != nil {
			return err
		}

		return base.ErrNotFound
	}

	return nil
}

func bookingUpdateQuery(book dbmodel.Booking) sq.UpdateBuilder {
	return base.Builder().
		Update(bookingTableName).
		SetMap(map[string]interface{}{
			"updated_at":        book.UpdatedAt,
			"slot":              book.Slot,
			"goal":              book.Goal,
			"title":             book.Title,
			"city":              book.City,
			"timeline_start_at": book.TimelineStartAt,
			"timeline_end_at":   book.TimelineEndAt,
			"description":       book.Description,
			"type":              book.Type,
			"links":             book.Links,
		}).
		Where("id=?", book.ID)
}

func bookingUpdateStatusQuery(book dbmodel.Booking) sq.UpdateBuilder {
	return base.Builder().
		Update(bookingTableName).
		SetMap(map[string]interface{}{
			"updated_at": book.UpdatedAt,
			"status":     book.Status,
		}).
		Where("id=?", book.ID)
}

func bookingUpdateStateQuery(book dbmodel.Booking) sq.UpdateBuilder {
	return base.Builder().
		Update(bookingTableName).
		SetMap(map[string]interface{}{
			"updated_at": book.UpdatedAt,
			"state":      book.State,
		}).
		Where("id=?", book.ID)
}

func bookingUpdateLinksQuery(book dbmodel.Booking) sq.UpdateBuilder {
	return base.Builder().
		Update(bookingTableName).
		SetMap(map[string]interface{}{
			"updated_at": book.UpdatedAt,
			"links":      book.Links,
		}).
		Where("id=?", book.ID)
}

func bookingDepartmentsDeleteQuery(book dbmodel.Booking) sq.DeleteBuilder {
	return base.Builder().
		Delete("booking_departments").
		Where("booking_id=?", book.ID)
}

func bookingDepartmentsInsertQuery(book dbmodel.Booking) sq.InsertBuilder {
	q := base.Builder().
		Insert("booking_departments").
		Columns("booking_id", "department_id")

	for _, department := range book.Departments {
		q = q.Values(book.ID, department.ID)
	}

	return q
}
