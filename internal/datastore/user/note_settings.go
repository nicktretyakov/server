package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"be/internal/datastore/base"
	"be/internal/model"
)

func (s Storage) GetNoteSettingByUser(ctx context.Context,
	userID *uuid.UUID,
) (*model.NoteSettings, error) {
	var setting model.NoteSettings

	err := s.db.Get(ctx, selectNoteSettingQuery(userID), &setting)
	if err != nil {
		return nil, err
	}

	return &setting, nil
}

func (s Storage) SetEmailNoteSetting(ctx context.Context,
	userID *uuid.UUID, emailOn bool,
) error {
	sql := updateNoteSettingsQuery("email_on", emailOn, userID)

	cmdTag, err := s.db.ExecBuilder(ctx, sql)
	if err != nil || cmdTag.RowsAffected() == 0 {
		if err != nil {
			return err
		}

		return base.ErrNotFound
	}

	return nil
}

func (s Storage) SetLifeNoteSetting(ctx context.Context,
	userID *uuid.UUID, lifeOn bool,
) error {
	sql := updateNoteSettingsQuery("life_on", lifeOn, userID)

	cmdTag, err := s.db.ExecBuilder(ctx, sql)
	if err != nil || cmdTag.RowsAffected() == 0 {
		if err != nil {
			return err
		}

		return base.ErrNotFound
	}

	return nil
}

func (s Storage) CreateNoteSettings(ctx context.Context, userID *uuid.UUID,
	emailOn, lifeOn bool,
) (*model.NoteSettings, error) {
	settings := model.NoteSettings{
		UserID:  userID,
		EmailOn: emailOn,
		LifeOn:  lifeOn,
	}

	sql := insertNoteSettingsQuery(settings)

	if _, err := s.db.ExecBuilder(ctx, sql); err != nil {
		return nil, err
	}

	return &settings, nil
}

func insertNoteSettingsQuery(u model.NoteSettings) sq.InsertBuilder {
	return base.Builder().
		Insert(noteSettingsTableName).
		Columns(
			"user_id",
			"email_on",
			"life_on",
		).
		Values(
			u.UserID,
			u.EmailOn,
			u.LifeOn,
		)
}

func updateNoteSettingsQuery(field string, value bool, userID *uuid.UUID) sq.UpdateBuilder {
	return base.Builder().
		Update(noteSettingsTableName).
		SetMap(map[string]interface{}{
			field: value,
		}).Where("user_id=?", userID)
}

func selectNoteSettingQuery(userID *uuid.UUID) sq.SelectBuilder {
	return base.Builder().
		Select(
			"user_id",
			"email_on",
			"life_on",
		).
		From(noteSettingsTableName).
		Where("user_id=?", userID)
}
