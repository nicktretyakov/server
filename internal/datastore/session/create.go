package session

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"be/internal/datastore/base"
	"be/internal/lib"
	"be/internal/model"
)

func (s storage) NewSession(ctx context.Context, session model.Session) (*model.Session, error) {
	if session.ID == uuid.Nil {
		session.ID = lib.UUID()
	}

	sql := sessionInsertQuery(session)

	if _, err := s.db.ExecBuilder(ctx, sql); err != nil {
		return nil, err
	}

	return &session, nil
}

func sessionInsertQuery(session model.Session) sq.InsertBuilder {
	return base.Builder().
		Insert(sessionTableName).
		Columns(
			"id",
			"created_at",
			"user_id",
			"access_token",
			"refresh_token",
			"refresh_expires",
			"profile_access_token",
			"profile_refresh_token",
		).
		Values(
			session.ID,                  // id
			session.CreatedAt,           // created_at
			session.UserID,              // user_id
			session.AccessToken,         // access_token
			session.RefreshToken,        // refresh_token
			session.RefreshExpires,      // refresh_expires
			session.ProfileAccessToken,  // profile_access_token
			session.ProfileRefreshToken, // profile_refresh_token
		)
}
