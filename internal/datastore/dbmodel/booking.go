package dbmodel

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"

	"be/internal/model"
)

type Booking struct {
	ID              uuid.UUID          `db:"id" yaml:"id"`
	Number          uint64             `db:"number" yaml:"number"`
	CreatedAt       time.Time          `db:"created_at" yaml:"created_at"`
	UpdatedAt       time.Time          `db:"updated_at" yaml:"updated_at"`
	Type            model.BookingType  `db:"type" yaml:"type"`
	Title           string             `db:"title" yaml:"title"`
	City            string             `db:"city" yaml:"city"`
	TimelineStartAt time.Time          `db:"timeline_start_at" yaml:"timeline_start_at"`
	TimelineEndAt   time.Time          `db:"timeline_end_at" yaml:"timeline_end_at"`
	Description     string             `db:"description" yaml:"description"`
	Departments     []model.Department `db:"departments" yaml:"-"`
	Reports         ReportBookingList  `db:"-" yaml:"-"`
	Outmembers      []model.Outmember  `db:"-" yaml:"-"`
	FinalReport     FinalReportOmit    `db:"final_report" yaml:"-"`
	Slot            pgtype.Numeric     `db:"slot" yaml:"slot"`
	Supervisor      User               `db:"supervisor" yaml:"-"`
	Author          User               `db:"author" yaml:"-"`
	Assignee        User               `db:"assignee" yaml:"-"`
	Participants    ParticipantList    `db:"participants" yaml:"-"`
	Goal            string             `db:"goal" yaml:"goal"`
	Status          model.Status       `db:"status" yaml:"status"`
	Attachments     AttachmentList     `db:"-"`
	Links           []model.Link       `db:"links"`
	State           model.State        `db:"state"`
}

func (p Booking) ToModel() model.Booking {
	book := model.Booking{
		ID:        p.ID,
		Number:    p.Number,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		Type:      p.Type,
		Title:     p.Title,
		City:      p.City,
		Timeline: model.Timeline{
			StartAt: p.TimelineStartAt,
			EndAt:   p.TimelineEndAt,
		},
		Description:  p.Description,
		Departments:  p.Departments,
		Reports:      p.Reports.ToModel(),
		FinalReport:  p.FinalReport.ToModel(),
		Slot:         ToNotification(p.Slot),
		Supervisor:   nil,
		Assignee:     nil,
		Author:       nil,
		Goal:         p.Goal,
		Status:       p.Status,
		Outmembers:   p.Outmembers,
		Attachments:  p.Attachments.Attachments(),
		Links:        p.Links,
		State:        p.State,
		Participants: p.Participants.Participants(),
	}

	if p.Supervisor.ID != nil {
		tmp := p.Supervisor.ToModel()
		book.Supervisor = &tmp
	}

	if p.Assignee.ID != nil {
		tmp := p.Assignee.ToModel()
		book.Assignee = &tmp
	}

	if p.Author.ID != nil {
		tmp := p.Author.ToModel()
		book.Author = &tmp
	}

	return book
}

func (p Booking) ToModelPtr() *model.Booking {
	pPtr := p.ToModel()
	return &pPtr
}

func (p *Booking) SetUsers(users []*BookingUser) {
	participants := make(ParticipantList, 0, len(users))

	for _, bookingUser := range users {
		switch bookingUser.Role {
		case model.SupervisorBookingUserRole:
			p.Supervisor = bookingUser.User
		case model.AssigneeBookingUserRole:
			p.Assignee = bookingUser.User
		case model.AuthorBookingUserRole:
			p.Author = bookingUser.User
		case model.ParticipantBookingUserRole:
			if bookingUser.User.ID == nil {
				continue
			}

			var roleTitle string

			if bookingUser.RoleTitle != nil {
				roleTitle = *bookingUser.RoleTitle
			}

			participants = append(participants, Participant{
				UserID: *bookingUser.User.ID,
				User:   &bookingUser.User,
				Role:   roleTitle,
			})
		}
	}

	p.Participants = participants
}

func BookingFromModel(p model.Booking) Booking {
	slot := pgtype.Numeric{}
	_ = slot.Set(p.Slot.String())

	dbBooking := Booking{
		ID:              p.ID,
		Number:          p.Number,
		CreatedAt:       p.CreatedAt,
		UpdatedAt:       p.UpdatedAt,
		Type:            p.Type,
		Title:           p.Title,
		City:            p.City,
		TimelineStartAt: p.Timeline.StartAt,
		TimelineEndAt:   p.Timeline.EndAt,
		Description:     p.Description,
		Departments:     p.Departments,
		Reports:         ReportListFromModel(p.Reports),
		Outmembers:      p.Outmembers,
		FinalReport:     FinalReportOmitFromModel(p.FinalReport),
		Slot:            slot,
		Goal:            p.Goal,
		Status:          p.Status,
		Attachments:     AttachmentListFromModel(p.Attachments),
		Links:           p.Links,
		State:           p.State,
		Participants:    ParticipantsFromModel(p.Participants),
	}

	if p.Supervisor != nil {
		dbBooking.Supervisor = UserFromModel(*p.Supervisor)
	}

	if p.Assignee != nil {
		dbBooking.Assignee = UserFromModel(*p.Assignee)
	}

	if p.Author != nil {
		dbBooking.Author = UserFromModel(*p.Author)
	}

	return dbBooking
}

type BookingList []*Booking

func (l BookingList) Bookings() []model.Booking {
	modelsList := make([]model.Booking, 0, len(l))
	for _, book := range l {
		modelsList = append(modelsList, book.ToModel())
	}

	return modelsList
}
