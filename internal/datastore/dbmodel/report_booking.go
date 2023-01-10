package dbmodel

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"

	"be/internal/model"
)

type ReportBookingList []ReportBooking

func (l ReportBookingList) ToModel() []model.ReportBooking {
	modelsList := make([]model.ReportBooking, 0, len(l))

	for _, r := range l {
		modelsList = append(modelsList, r.ToModel())
	}

	return modelsList
}

func ReportListFromModel(reports []model.ReportBooking) ReportBookingList {
	repList := make(ReportBookingList, 0, len(reports))

	for _, report := range reports {
		repList = append(repList, ReportFromModel(report))
	}

	return repList
}

type ReportBooking struct {
	ID        uuid.UUID          `db:"id" yaml:"id"`
	CreatedAt time.Time          `db:"created_at" yaml:"created_at"`
	UpdatedAt time.Time          `db:"updated_at" yaml:"updated_at"`
	BookingID uuid.UUID          `db:"booking_id" yaml:"booking_id"`
	Period    time.Time          `db:"period" yaml:"period"`
	Slot    pgtype.Numeric     `db:"slot" yaml:"-"`
	EndAt     *time.Time         `db:"end_at" yaml:"end_at"`
	Events    *string            `db:"events" yaml:"events"`
	Reasons   *string            `db:"reasons" yaml:"reasons"`
	Comment   *string            `db:"comment" yaml:"comment"`
	Status    model.ReportStatus `db:"status" yaml:"status"`
}

func (r ReportBooking) ToModel() model.ReportBooking {
	return model.ReportBooking{
		ID:        r.ID,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
		BookingID: r.BookingID,
		Period: model.Period{
			Year:  int32(r.Period.Year()),
			Month: int32(r.Period.Month()),
		},
		Slot:  ToNotificationPtr(r.Slot),
		EndAt:   r.EndAt,
		Events:  r.Events,
		Reasons: r.Reasons,
		Comment: r.Comment,
		Status:  r.Status,
	}
}

func (r ReportBooking) ToModelPtr() *model.ReportBooking {
	repPtr := r.ToModel()
	return &repPtr
}

func ReportFromModel(p model.ReportBooking) ReportBooking {
	var slot pgtype.Numeric

	_ = slot.Set(nil)

	if p.Slot != nil {
		_ = slot.Set(p.Slot.String())
	}

	return ReportBooking{
		ID:        p.ID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		BookingID: p.BookingID,
		Period:    p.Period.Time(),
		Slot:    slot,
		EndAt:     p.EndAt,
		Events:    p.Events,
		Reasons:   p.Reasons,
		Comment:   p.Comment,
		Status:    p.Status,
	}
}
