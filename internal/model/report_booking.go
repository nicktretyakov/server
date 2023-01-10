package model

import (
	"time"

	"github.com/google/uuid"
)

type ReportStatus int

const (
	DeletedReportStatus ReportStatus = 0
	NotSendReportStatus ReportStatus = 1
	SendReportStatus    ReportStatus = 2
)

type ReportBooking struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	BookingID uuid.UUID
	Period    Period
	Slot    *Notification
	EndAt     *time.Time
	Events    *string
	Reasons   *string
	Comment   *string
	Status    ReportStatus
}

func (r ReportBooking) GetSlot() Notification {
	if r.Slot == nil {
		return Notification{}
	}

	return *r.Slot
}

func (r ReportBooking) GetEndAt() time.Time {
	if r.EndAt == nil {
		return time.Time{}
	}

	return *r.EndAt
}

func (r ReportBooking) GetEvents() string {
	if r.Events == nil {
		return ""
	}

	return *r.Events
}

func (r ReportBooking) GetReasons() string {
	if r.Reasons == nil {
		return ""
	}

	return *r.Reasons
}

func (r ReportBooking) GetComment() string {
	if r.Comment == nil {
		return ""
	}

	return *r.Comment
}

func (r ReportBooking) NotRelevant() bool {
	return r.Period.Time().Before(time.Now()) && r.Status == NotSendReportStatus
}
