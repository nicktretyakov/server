package report

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"be/internal/datastore/base"
	"be/internal/datastore/dbmodel"
	"be/internal/model"
)

func (s *Storage) FindBookingReportByID(ctx context.Context, reportID uuid.UUID) (*model.ReportBooking, error) {
	query := SelectQuery().Where("reports.id=?", reportID)

	var rep dbmodel.ReportBooking
	if err := s.db.Get(ctx, query, &rep); err != nil {
		return nil, err
	}

	return rep.ToModelPtr(), nil
}

func (s *Storage) FindRoomReportByID(ctx context.Context, reportID uuid.UUID) (*model.ReportRoom, error) {
	var rep dbmodel.ReportRoom

	if err := s.db.Get(ctx, reportsRoomSelectQuery(reportID), &rep); err != nil {
		return nil, err
	}

	return rep.ToModelPtr(), nil
}

func reportsRoomSelectQuery(reportID uuid.UUID) sq.SelectBuilder {
	return base.
		Builder().
		Select(
			"id",
			"created_at",
			"updated_at",
			"room_id",
			"timeline_start_at",
			"timeline_end_at",
			"comment",
			"slots",
			"equipments",
			"releases",
		).
		From(reportRoomTableName).
		Where("id = ?", reportID)
}
