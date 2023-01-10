package address

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"be/internal/datastore/base"
	"be/internal/lib"
	"be/internal/model"
)

func (s *Storage) LifeNotesList(ctx context.Context) ([]*model.LifeNote, error) {
	notes := make([]*model.LifeNote, 0)

	if err := s.db.Select(ctx, lifeNoteSelectQuery(), &notes); err != nil {
		return nil, err
	}

	return notes, nil
}

func (s *Storage) CreateLifeNote(ctx context.Context, note model.LifeNote) (*model.LifeNote, error) {
	note.ID = lib.UUID()
	note.CreatedAt = s.db.Now()

	if _, err := s.db.ExecBuilder(ctx, lifeNoteInsertQuery(note)); err != nil {
		return nil, err
	}

	return &note, nil
}

func (s *Storage) UpdateLifeNote(ctx context.Context, note model.LifeNote) (*model.LifeNote, error) {
	note.SentAt = s.db.Now()

	if _, err := s.db.ExecBuilder(ctx, lifeNoteUpdateQuery(note)); err != nil {
		return nil, err
	}

	return &note, nil
}

func (s *Storage) GetChannelFromLifeNote(ctx context.Context, userID *uuid.UUID) (uuid.UUID, error) {
	var channelID uuid.UUID

	if err := s.db.Get(ctx, lifeNoteGetQuery(*userID), &channelID); err != nil {
		return uuid.Nil, err
	}

	return channelID, nil
}

func lifeNoteSelectQuery() sq.SelectBuilder {
	return base.Builder().
		Select(
			"id",
			"event",
			"actor_id",
			"object",
			"recipient_id",
			"channel_id",
			"bot_id",
			"body",
			"status",
			"for_entities",
			"created_at",
			"sent_at",
		).
		From(lifeNoteTableName).
		OrderBy("created_at desc").
		Where("status=?", model.NotSend)
}

func lifeNoteInsertQuery(note model.LifeNote) sq.InsertBuilder {
	return base.Builder().
		Insert(lifeNoteTableName).
		Columns(
			"id",
			"event",
			"actor_id",
			"object",
			"recipient_id",
			"channel_id",
			"bot_id",
			"body",
			"status",
			"for_entities",
			"created_at",
			"sent_at",
		).
		Values(
			note.ID,
			note.Event,
			note.ActorID,
			note.Object,
			note.RecipientID,
			note.ChannelID,
			note.BotID,
			note.Body,
			note.Status,
			note.ForEntities,
			note.CreatedAt,
			note.SentAt,
		)
}

func lifeNoteUpdateQuery(note model.LifeNote) sq.UpdateBuilder {
	return base.Builder().
		Update(lifeNoteTableName).
		SetMap(map[string]interface{}{
			"channel_id": note.ChannelID,
			"sent_at":    note.SentAt,
			"status":     note.Status,
		}).
		Where("id=?", note.ID)
}

func lifeNoteGetQuery(userID uuid.UUID) sq.SelectBuilder {
	return base.Builder().
		Select("channel_id").
		From(lifeNoteTableName).
		Where("channel_id !=? AND recipient_id=?", uuid.Nil, userID).
		GroupBy("channel_id")
}
