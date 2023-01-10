package stage

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"be/internal/datastore/base"
	"be/internal/datastore/dbmodel"
)

func selectQuery() sq.SelectBuilder {
	return base.Builder().
		Select("stages.id",
			"stages.booking_id",
			"stages.title",
			"stages.start_at",
			"stages.end_at",
			"stages.status",
			"stages.created_at").
		From(stageTableNameAlias).
		OrderBy("created_at desc")
}

func insertQuery(item dbmodel.Stage) sq.InsertBuilder {
	return base.Builder().
		Insert(stageTableName).
		Columns(
			"id",
			"booking_id",
			"title",
			"start_at",
			"end_at",
			"created_at",
			"status").
		Values(
			item.ID,
			item.BookingID,
			item.Title,
			item.StartAt,
			item.EndAt,
			item.CreatedAt,
			item.Status,
		)
}

func stageUpdateQuery(item dbmodel.Stage) sq.UpdateBuilder {
	return base.Builder().
		Update(stageTableName).
		SetMap(map[string]interface{}{
			"title":    item.Title,
			"start_at": item.StartAt,
			"end_at":   item.EndAt,
			"status":   item.Status,
		}).
		Where("id=?", item.ID)
}

func deleteQuery(stageID uuid.UUID) sq.DeleteBuilder {
	return base.Builder().Delete(stageTableName).Where("id=?", stageID)
}
