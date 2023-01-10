package model

import (
	"time"

	"github.com/google/uuid"
)

type FinalReportStatus int

const (
	DeletedFinalReportStatus    FinalReportStatus = 0
	InitialFinalReportStatus    FinalReportStatus = 1
	DeclinedFinalReportStatus   FinalReportStatus = 2
	OnRegisterFinalReportStatus FinalReportStatus = 3
	ConfirmedFinalReportStatus  FinalReportStatus = 4
	OnAgreeFinalReportStatus    FinalReportStatus = 5
)

func (s FinalReportStatus) IsOnAgree() bool {
	return s.Eq(OnAgreeFinalReportStatus)
}

func (s FinalReportStatus) IsDeclined() bool {
	return s.Eq(DeclinedFinalReportStatus)
}

func (s FinalReportStatus) Eq(s2 FinalReportStatus) bool {
	return s == s2
}

func (s FinalReportStatus) In(s2 ...FinalReportStatus) bool {
	for _, status := range s2 {
		if s.Eq(status) {
			return true
		}
	}

	return false
}

type FinalReport struct {
	ID              uuid.UUID         `yaml:"id"`
	CreatedAt       time.Time         `yaml:"created_at"`
	UpdatedAt       time.Time         `yaml:"updated_at"`
	BookingID       uuid.UUID         `yaml:"booking_id"`
	Slot          Notification             `yaml:"slot"`
	EndAt           time.Time         `yaml:"end_at"`
	Comment         string            `yaml:"comment"`
	Status          FinalReportStatus `yaml:"status"`
	AttachmentsUUID []uuid.UUID       `yaml:"-"`
}
