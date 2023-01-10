package address

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"be/internal/datastore/base"
	"be/internal/lib"
	"be/internal/model"
)

func (s *Storage) CreateEmailNote(ctx context.Context, note model.EmailNote) (*model.EmailNote, error) {
	note.ID = lib.UUID()
	note.CreatedAt = s.db.Now()

	if _, err := s.db.ExecBuilder(ctx, noteInsertQuery(note)); err != nil {
		return nil, err
	}

	return &note, nil
}

func (s *Storage) UpdateEmailNote(ctx context.Context, note model.EmailNote) (*model.EmailNote, error) {
	note.SentAt = s.db.Now()

	if _, err := s.db.ExecBuilder(ctx, noteUpdateQuery(note)); err != nil {
		return nil, err
	}

	return &note, nil
}

func (s *Storage) EmailNotesList(ctx context.Context) ([]*model.EmailNote, error) {
	notes := make([]*model.EmailNote, 0)

	if err := s.db.Select(ctx, noteSelectQuery(), &notes); err != nil {
		return nil, err
	}

	return notes, nil
}

func noteInsertQuery(note model.EmailNote) sq.InsertBuilder {
	return base.Builder().
		Insert(emailNoteTableName).
		Columns(
			"id",
			"event",
			"actor_id",
			"object",
			"recipient_id",
			"recipient_email",
			"sender_email",
			"subject",
			"body",
			"status",
			"created_at",
			"sent_at",
		).
		Values(
			note.ID,
			note.Event,
			note.ActorID,
			note.Object,
			note.RecipientID,
			note.RecipientEmail,
			note.SenderEmail,
			note.Subject,
			note.Body,
			note.Status,
			note.CreatedAt,
			note.SentAt,
		)
}

func noteUpdateQuery(note model.EmailNote) sq.UpdateBuilder {
	return base.Builder().
		Update(emailNoteTableName).
		SetMap(map[string]interface{}{
			"sent_at": note.SentAt,
			"status":  note.Status,
		}).
		Where("id=?", note.ID)
}

func noteSelectQuery() sq.SelectBuilder {
	return base.Builder().
		Select(
			"id",
			"event",
			"actor_id",
			"object",
			"recipient_id",
			"recipient_email",
			"sender_email",
			"subject",
			"body",
			"status",
			"created_at",
			"sent_at",
		).
		From(emailNoteTableName).
		OrderBy("created_at desc").
		Where("status=?", model.NotSend)
}
