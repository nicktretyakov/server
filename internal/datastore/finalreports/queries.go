package finalreports

import (
	sq "github.com/Masterminds/squirrel"

	"be/internal/datastore/base"
	"be/internal/datastore/dbmodel"
)

func SelectQuery() sq.SelectBuilder {
	return base.
		Builder().
		Select(
			"final_reports.id",
			"final_reports.created_at",
			"final_reports.updated_at",
			"final_reports.booking_id",
			"final_reports.slot",
			"final_reports.end_at",
			"final_reports.comment",
			"final_reports.status",
			"final_reports.attachments_uuid",
		).
		From(finalReportsTableNameAlias)
}

func reportUpdateQuery(rep dbmodel.FinalReport) sq.UpdateBuilder {
	return base.Builder().
		Update(finalReportsTableName).
		SetMap(map[string]interface{}{
			"updated_at":       rep.UpdatedAt,
			"slot":           rep.Slot,
			"end_at":           rep.EndAt,
			"comment":          rep.Comment,
			"status":           rep.Status,
			"attachments_uuid": rep.AttachmentsUUID,
		}).
		Where("id=?", rep.ID)
}

func reportStatusUpdateQuery(rep dbmodel.FinalReport) sq.UpdateBuilder {
	return base.Builder().
		Update(finalReportsTableName).
		SetMap(map[string]interface{}{
			"updated_at": rep.UpdatedAt,
			"status":     rep.Status,
		}).
		Where("booking_id=?", rep.BookingID)
}
