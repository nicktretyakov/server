package report

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"be/internal/datastore/base"
	"be/internal/datastore/dbmodel"
	"be/internal/lib"
	"be/internal/model"
)

func (s *Storage) BulkCreate(ctx context.Context, reports []model.ReportBooking) ([]model.ReportBooking, error) {
	now := s.db.Now()
	dbReports := make(dbmodel.ReportBookingList, 0, len(reports))

	for _, report := range reports {
		dbRep := dbmodel.ReportFromModel(report)
		dbRep.ID = lib.UUID()

		dbRep.CreatedAt = now
		dbRep.UpdatedAt = now

		dbReports = append(dbReports, dbRep)
	}

	if _, err := s.db.ExecBuilder(ctx, insertQuery(dbReports)); err != nil {
		return nil, err
	}

	return dbReports.ToModel(), nil
}

func (s *Storage) CreateReportRoom(ctx context.Context, report model.ReportRoom) (*model.ReportRoom, error) {
	dbReport := dbmodel.ReportRoomFromModel(report)
	dbReport.ID = lib.UUID()

	if _, err := s.db.ExecBuilder(ctx, insertRoomReportQuery(dbReport)); err != nil {
		return nil, err
	}

	return dbReport.ToModelPtr(), nil
}

func insertQuery(reports dbmodel.ReportBookingList) sq.InsertBuilder {
	q := base.Builder().
		Insert(reportBookingTableName).
		Columns(
			"id",
			"created_at",
			"updated_at",
			"booking_id",
			"period",
			"slot",
			"end_at",
			"events",
			"reasons",
			"comment",
			"status",
		)

	for _, rep := range reports {
		q = q.Values(
			rep.ID,
			rep.CreatedAt,
			rep.UpdatedAt,
			rep.BookingID,
			rep.Period,
			rep.Slot,
			rep.EndAt,
			rep.Events,
			rep.Reasons,
			rep.Comment,
			rep.Status,
		)
	}

	return q
}

func insertRoomReportQuery(report dbmodel.ReportRoom) sq.InsertBuilder {
	return base.Builder().
		Insert(reportRoomTableName).
		Columns(
			"id",
			"room_id",
			"timeline_start_at",
			"timeline_end_at",
			"comment",
			"slots",
			"equipments",
			"releases",
		).
		Values(
			report.ID,
			report.RoomID,
			report.TimelineStartAt,
			report.TimelineEndAt,
			report.Comment,
			report.Slots,
			report.Equipments,
			report.Releases,
		)
}
