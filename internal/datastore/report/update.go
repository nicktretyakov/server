package report

import (
	"context"

	"be/internal/datastore/base"
	"be/internal/datastore/dbmodel"
	"be/internal/model"
)

func (s *Storage) Update(ctx context.Context, rep model.ReportBooking) (*model.ReportBooking, error) {
	repToStore := dbmodel.ReportFromModel(rep)
	repToStore.UpdatedAt = s.db.Now()

	sql := reportUpdateQuery(repToStore)

	if cmd, err := s.db.ExecBuilder(ctx, sql); err != nil || cmd.RowsAffected() == 0 {
		if err != nil {
			return nil, err
		}

		return nil, base.ErrNotFound
	}

	return repToStore.ToModelPtr(), nil
}
