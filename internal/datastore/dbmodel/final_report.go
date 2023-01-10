package dbmodel

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"

	"be/internal/model"
)

type FinalReport struct {
	ID              uuid.UUID               `db:"id"`
	CreatedAt       time.Time               `db:"created_at"`
	UpdatedAt       time.Time               `db:"updated_at"`
	BookingID       uuid.UUID               `db:"booking_id"`
	Slot          pgtype.Numeric          `db:"slot"`
	EndAt           time.Time               `db:"end_at"`
	Comment         string                  `db:"comment"`
	Status          model.FinalReportStatus `db:"status"`
	AttachmentsUUID []uuid.UUID             `db:"attachments_uuid"`
}

func (r FinalReport) ToModel() model.FinalReport {
	return model.FinalReport{
		ID:              r.ID,
		CreatedAt:       r.CreatedAt,
		UpdatedAt:       r.UpdatedAt,
		BookingID:       r.BookingID,
		Slot:          ToNotification(r.Slot),
		EndAt:           r.EndAt,
		Comment:         r.Comment,
		Status:          r.Status,
		AttachmentsUUID: r.AttachmentsUUID,
	}
}

func (r FinalReport) ToModelPtr() *model.FinalReport {
	repPtr := r.ToModel()
	return &repPtr
}

func FinalReportFromModel(p model.FinalReport) FinalReport {
	var slot pgtype.Numeric

	_ = slot.Set(p.Slot.String())

	return FinalReport{
		ID:              p.ID,
		CreatedAt:       p.CreatedAt,
		UpdatedAt:       p.UpdatedAt,
		BookingID:       p.BookingID,
		Slot:          slot,
		EndAt:           p.EndAt,
		Comment:         p.Comment,
		Status:          p.Status,
		AttachmentsUUID: p.AttachmentsUUID,
	}
}

func FinalReportOmitFromModel(p model.FinalReport) FinalReportOmit {
	var slot pgtype.Numeric

	_ = slot.Set(p.Slot.String())

	return FinalReportOmit{
		ID:              &p.ID,
		CreatedAt:       &p.CreatedAt,
		UpdatedAt:       &p.UpdatedAt,
		BookingID:       &p.BookingID,
		Slot:          slot,
		EndAt:           &p.EndAt,
		Comment:         &p.Comment,
		Status:          &p.Status,
		AttachmentsUUID: p.AttachmentsUUID,
	}
}

type FinalReportOmit struct {
	ID              *uuid.UUID               `db:"id"`
	CreatedAt       *time.Time               `db:"created_at"`
	UpdatedAt       *time.Time               `db:"updated_at"`
	BookingID       *uuid.UUID               `db:"booking_id"`
	Slot          pgtype.Numeric           `db:"slot"`
	EndAt           *time.Time               `db:"end_at"`
	Comment         *string                  `db:"comment"`
	Status          *model.FinalReportStatus `db:"status"`
	AttachmentsUUID []uuid.UUID              `db:"attachments_uuid"`
}

func (o FinalReportOmit) ToModel() model.FinalReport {
	if o.ID == nil {
		return model.FinalReport{}
	}

	return FinalReport{
		ID:              *o.ID,
		CreatedAt:       *o.CreatedAt,
		UpdatedAt:       *o.UpdatedAt,
		BookingID:       *o.BookingID,
		Slot:          o.Slot,
		EndAt:           *o.EndAt,
		Comment:         *o.Comment,
		Status:          *o.Status,
		AttachmentsUUID: o.AttachmentsUUID,
	}.ToModel()
}
