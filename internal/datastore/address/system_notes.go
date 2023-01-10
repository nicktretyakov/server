package address

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"be/internal/datastore/base"
	"be/internal/datastore/sorting"
	"be/internal/lib"
	"be/internal/model"
)

func (s *Storage) SystemNoteList(
	ctx context.Context,
	userID uuid.UUID,
	status model.SystemStatus,
	limit,
	offset uint32,
	sorting sorting.Sorting,
) ([]*model.SystemNote, error) {
	notes := make([]*model.SystemNote, 0)

	query := systemNoteSelectQuery(userID, status, limit, offset)
	query = sorting.Apply(query)

	if err := s.db.Select(ctx, query, &notes); err != nil {
		return nil, err
	}

	return notes, nil
}

func (s *Storage) GetSystemNotesCount(ctx context.Context, userID uuid.UUID) (int, error) {
	var count int

	if err := s.db.Get(ctx, systemNotesCountSelectQuery(userID), &count); err != nil {
		return 0, err
	}

	return count, nil
}

func (s *Storage) UpdateSystemNotes(ctx context.Context, noteIDs []uuid.UUID) error {
	if cmd, err := s.db.ExecBuilder(ctx, systemNotesUpdateQuery(noteIDs)); err != nil || cmd.RowsAffected() == 0 {
		if err != nil {
			return err
		}

		return base.ErrNotFound
	}

	return nil
}

func (s *Storage) CreateSystemNote(ctx context.Context, note model.SystemNote) (*model.SystemNote, error) {
	note.ID = lib.UUID()
	note.CreatedAt = s.db.Now()

	if _, err := s.db.ExecBuilder(ctx, systemNoteInsertQuery(note)); err != nil {
		return nil, err
	}

	return &note, nil
}

func systemNoteSelectQuery(userID uuid.UUID, status model.SystemStatus, limit, offset uint32) sq.SelectBuilder {
	return base.Builder().
		Select(
			"id",
			"event",
			"actor_id",
			"address_id",
			"object",
			"recipient_id",
			"header",
			"body",
			"status",
			"created_at",
			"read_at",
		).
		From(systemNoteTableName).
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		Where("status=? and recipient_id=?", status, userID)
}

func systemNotesCountSelectQuery(userID uuid.UUID) sq.SelectBuilder {
	return base.Builder().
		Select("count(*)").
		From(systemNoteTableName).
		Where("status=? and recipient_id=?", model.NotRead, userID)
}

func systemNotesUpdateQuery(noteIDs []uuid.UUID) sq.UpdateBuilder {
	return base.Builder().
		Update(systemNoteTableName).
		SetMap(map[string]interface{}{
			"read_at": time.Now(),
			"status":  model.Read,
		}).
		Where(sq.Eq{"id": noteIDs})
}

func systemNoteInsertQuery(note model.SystemNote) sq.InsertBuilder {
	return base.Builder().
		Insert(systemNoteTableName).
		Columns(
			"id",
			"event",
			"actor_id",
			"address_id",
			"object",
			"recipient_id",
			"header",
			"body",
			"status",
			"created_at",
			"read_at",
		).
		Values(
			note.ID,
			note.Event,
			note.ActorID,
			note.AddressID,
			note.Object,
			note.RecipientID,
			note.Header,
			note.Body,
			note.Status,
			note.CreatedAt,
			note.ReadAt,
		)
}
