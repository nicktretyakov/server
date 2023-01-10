package testutil

import (
	"time"

	"github.com/google/uuid"

	"be/internal/model"
)

type booking struct {
	ID              uuid.UUID         `yaml:"id"`
	Number          uint64            `yaml:"number"`
	CreatedAt       time.Time         `yaml:"created_at"`
	UpdatedAt       time.Time         `yaml:"updated_at"`
	TimelineStartAt time.Time         `yaml:"timeline_start_at"`
	TimelineEndAt   time.Time         `yaml:"timeline_end_at"`
	Title           string            `yaml:"title"`
	City            string            `yaml:"city"`
	Description     string            `yaml:"description"`
	Slot            float64           `yaml:"slot"`
	Goal            string            `yaml:"goal"`
	Status          model.Status      `yaml:"status"`
	Type            model.BookingType `yaml:"type"`
	Links           []model.Link      `yaml:"links"`
	State           model.State       `yaml:"state"`
}

type room struct {
	ID             uuid.UUID    `yaml:"id"`              // Идентификатор продукта
	Number         uint64       `yaml:"number"`          // Номер проекта
	Title          string       `yaml:"title"`           // Имя проекта
	Description    string       `yaml:"description"`     // Описание проекта
	TargetAudience string       `yaml:"target_audience"` // Целевая аудитория
	Status         model.Status `yaml:"status"`          // Статус агрегата
	Links          []model.Link `yaml:"links"`
	CreationDate   time.Time    `yaml:"creation_date"` // Дата создания продукта
	CreatedAt      time.Time    `yaml:"created_at"`    // Дата создания агрегата
	UpdatedAt      time.Time    `yaml:"updated_at"`    // Дата обновления агрегата
	Bookings       []uuid.UUID  `yaml:"booking_ids"`   // Проекты
	State          model.State  `yaml:"state"`
	Space          int32     `yaml:"space"`
	SecurityEmail  string    `yaml:"security_email"`
	Visible        bool    `yaml:"visible"`
}

func (p booking) toModel() model.Booking {
	return model.Booking{
		ID:          p.ID,
		Number:      p.Number,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		Type:        p.Type,
		Title:       p.Title,
		City:        p.City,
		Description: p.Description,
		Timeline: model.Timeline{
			StartAt: p.TimelineStartAt,
			EndAt:   p.TimelineEndAt,
		},
		Slot:         model.NewNotificationFromFloat(p.Slot),
		Goal:         p.Goal,
		Status:       p.Status,
		Links:        p.Links,
		State:        p.State,
		Participants: []model.Participant{},
	}
}

func (p room) toModel() model.Room {
	return model.Room{
		ID:             p.ID,
		Number:         p.Number,
		Title:          p.Title,
		Description:    p.Description,
		TargetAudience: p.TargetAudience,
		Status:         p.Status,
		Links:          p.Links,
		CreationDate:   p.CreationDate,
		CreatedAt:      p.CreatedAt,
		UpdatedAt:      p.UpdatedAt,
		State:          p.State,
		Space:           p.Space,
    	SecurityEmail:   p.SecurityEmail,
	    Visible:         p.Visible,
	}
}

type (
	bookingList []booking
	roomList    []room
)

func (l bookingList) Bookings() []model.Booking {
	modelsList := make([]model.Booking, 0, len(l))
	for _, book := range l {
		modelsList = append(modelsList, book.toModel())
	}

	return modelsList
}

func (l roomList) Rooms() []model.Room {
	modelsList := make([]model.Room, 0, len(l))
	for _, room := range l {
		modelsList = append(modelsList, room.toModel())
	}

	return modelsList
}

type bookingDepartment struct {
	BookingID    uuid.UUID `yaml:"booking_id"`
	DepartmentID uuid.UUID `yaml:"department_id"`
}
