package finalreports

import (
	"context"

	"github.com/google/uuid"

	"be/internal/datastore/base"
	"be/internal/datastore/dbmodel"
	"be/internal/model"
)

func (s *storage) UpdateStatus(ctx context.Context, bookingID uuid.UUID, status model.FinalReportStatus) error {
	repToStore := dbmodel.FinalReportFromModel(model.FinalReport{BookingID: bookingID, Status: status})
	repToStore.UpdatedAt = s.db.Now()

	sql := reportStatusUpdateQuery(repToStore)

	if cmd, err := s.db.ExecBuilder(ctx, sql); err != nil || cmd.RowsAffected() == 0 {
		if err != nil {
			return err
		}

		return base.ErrNotFound
	}

	return nil
}

func (s *storage) Update(ctx context.Context, rep model.FinalReport) (*model.FinalReport, error) {
	repToStore := dbmodel.FinalReportFromModel(rep)
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
