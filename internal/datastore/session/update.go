package session

import (
	sq "github.com/Masterminds/squirrel"

	"be/internal/datastore/base"
	"be/internal/model"
)

func sessionUpdateQuery(session model.Session) sq.UpdateBuilder {
	return base.Builder().
		Update(sessionTableName).
		SetMap(map[string]interface{}{
			"refresh_token":   session.RefreshToken,
			"refresh_expires": session.RefreshExpires,
			"access_token":    session.AccessToken,
		})
}
