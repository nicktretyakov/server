package dbmodel

import (
	"time"

	"github.com/google/uuid"

	"be/internal/model"
)

type (
	ReportRoom struct {
		ID              uuid.UUID       `db:"id" yaml:"id"`
		CreatedAt       time.Time       `db:"created_at" yaml:"created_at"`
		UpdatedAt       time.Time       `db:"updated_at" yaml:"updated_at"`
		RoomID       uuid.UUID       `db:"room_id" yaml:"room_id"`
		TimelineStartAt time.Time       `db:"timeline_start_at" yaml:"timeline_start_at"`
		TimelineEndAt   time.Time       `db:"timeline_end_at" yaml:"timeline_end_at"`
		Comment         *string         `db:"comment" yaml:"comment"`
		Slots         []model.Slot  `db:"slots" yaml:"slots"`
		Equipments         []model.Equipment  `db:"equipments" yaml:"equipments"`
		Releases        []model.Release `db:"releases" yaml:"releases"`
	}

	ReportRoomList []ReportRoom
)

func ReportRoomFromModel(r model.ReportRoom) ReportRoom {
	return ReportRoom{
		ID:              r.ID,
		CreatedAt:       r.CreatedAt,
		UpdatedAt:       r.UpdatedAt,
		RoomID:       r.RoomID,
		TimelineStartAt: r.Timeline.StartAt,
		TimelineEndAt:   r.Timeline.EndAt,
		Comment:         r.Comment,
		Slots:         r.Slots,
		Equipments:         r.Equipments,
		Releases:        r.Releases,
	}
}

func (r ReportRoom) ToModel() model.ReportRoom {
	return model.ReportRoom{
		ID:        r.ID,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
		RoomID: r.RoomID,
		Timeline: model.Timeline{
			StartAt: r.TimelineStartAt,
			EndAt:   r.TimelineEndAt,
		},
		Comment:  r.Comment,
		Equipments:  r.Equipments,
		Slots:  r.Slots,
		Releases: r.Releases,
	}
}

func (r ReportRoom) ToModelPtr() *model.ReportRoom {
	ptr := r.ToModel()
	return &ptr
}

func ToReportRoomList(reports []model.ReportRoom) ReportRoomList {
	reportList := make(ReportRoomList, 0, len(reports))
	for _, report := range reports {
		reportList = append(reportList, ReportRoomFromModel(report))
	}

	return reportList
}

func (r ReportRoomList) ReportsRoom() []model.ReportRoom {
	reports := make([]model.ReportRoom, 0, len(r))
	for _, report := range r {
		reports = append(reports, report.ToModel())
	}

	return reports
}
