package finalreports

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"be/internal/datastore/base"
	"be/internal/datastore/dbmodel"
	"be/internal/lib"
	"be/internal/model"
)

func (s *storage) Create(ctx context.Context, report model.FinalReport) (*model.FinalReport, error) {
	now := s.db.Now()
	dbReport := dbmodel.FinalReportFromModel(report)

	dbReport.ID = lib.UUID()
	dbReport.CreatedAt = now
	dbReport.UpdatedAt = now

	if _, err := s.db.ExecBuilder(ctx, insertQuery(dbReport)); err != nil {
		return nil, err
	}

	return dbReport.ToModelPtr(), nil
}

func insertQuery(rep dbmodel.FinalReport) sq.InsertBuilder {
	return base.Builder().
		Insert(finalReportsTableName).
		Columns(
			"id",
			"created_at",
			"updated_at",
			"booking_id",
			"slot",
			"end_at",
			"comment",
			"status",
			"attachments_uuid",
		).Values(
		rep.ID,
		rep.CreatedAt,
		rep.UpdatedAt,
		rep.BookingID,
		rep.Slot,
		rep.EndAt,
		rep.Comment,
		rep.Status,
		rep.AttachmentsUUID,
	)
}
