package session

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"be/internal/datastore/base"
	"be/internal/model"
)

func (s storage) FindSessionByRefreshToken(ctx context.Context, refresh string) (*model.Session, error) {
	return findOne(ctx, s.db, "refresh_token=? and refresh_expires>now()", refresh)
}

func (s storage) FindSessionByID(ctx context.Context, sessionID uuid.UUID) (*model.Session, error) {
	return findOne(ctx, s.db, "id=?", sessionID)
}

func findOne(ctx context.Context, db *base.DB, pred interface{}, args ...interface{}) (*model.Session, error) {
	var session model.Session

	err := db.Get(ctx, sessionSelectQuery().Where(pred, args...), &session)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

func sessionSelectQuery() sq.SelectBuilder {
	return base.Builder().
		Select(
			"id",
			"created_at",
			"user_id",
			"access_token",
			"refresh_token",
			"refresh_expires",
			"profile_access_token",
			"profile_refresh_token",
		).
		From(sessionTableName)
}
