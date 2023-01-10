package session

import (
	"context"

	"be/internal/datastore/base"
	"be/internal/model"
)

func (s storage) RefreshSession(ctx context.Context, session model.Session) error {
	sql := sessionUpdateQuery(session).Where("id=?", session.ID)

	cmdTag, err := s.db.ExecBuilder(ctx, sql)
	if err != nil || cmdTag.RowsAffected() == 0 {
		if err != nil {
			return err
		}

		return base.ErrNotFound
	}

	return nil
}
