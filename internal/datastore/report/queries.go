package report

import (
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"be/internal/datastore/base"
	"be/internal/datastore/dbmodel"
)

func SelectQuery() sq.SelectBuilder {
	return base.
		Builder().
		Select(
			"reports.id",
			"reports.created_at",
			"reports.updated_at",
			"reports.booking_id",
			"reports.period",
			"reports.slot",
			"reports.end_at",
			"reports.events",
			"reports.reasons",
			"reports.comment",
			"reports.status",
		).
		From(reportTableNameAlias)
}

func reportUpdateQuery(rep dbmodel.ReportBooking) sq.UpdateBuilder {
	return base.Builder().
		Update(reportBookingTableName).
		SetMap(map[string]interface{}{
			"updated_at": rep.UpdatedAt,
			"slot":     rep.Slot,
			"end_at":     rep.EndAt,
			"events":     rep.Events,
			"reasons":    rep.Reasons,
			"comment":    rep.Comment,
			"status":     rep.Status,
		}).
		Where("id=?", rep.ID)
}

func deleteQuery(reportID, bookingID uuid.UUID, period time.Time) sq.DeleteBuilder {
	return base.Builder().Delete(reportBookingTableName).Where("id=? and period=? and booking_id=?", reportID, period, bookingID)
}
