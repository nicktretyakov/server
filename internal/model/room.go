package model

import (
	"time"

	"github.com/google/uuid"
)

const (
	RoomColumnLowest  = "A"
	RoomColumnHighest = "E"

	ColumnNumberRoom       = "Номер"
	ColumnTitleRoom        = "Название"
	ColumnCreationDateRoom = "Фактическая дата создания"
	ColumnOwnerRoom        = "Владелец "
	ColumnStatusRoom       = "Статус "

	RoomPrefixOutputFilename = "Rooms"
)

type (
	Room struct {
		ID             uuid.UUID       `yaml:"id"`              // Идентификатор 
		Number         uint64          `yaml:"number"`          // Номер 
		Author         *User           `yaml:"author_id"`       // Автор продукта
		Owner          *User           `yaml:"owner_id"`        // Владелец продукта
		Employee       *User           `yaml:"employee_id"`     // Согласующий
		Title          string          `yaml:"title"`           // Имя 
		Description    string          `yaml:"description"`     // Описание
		TargetAudience string          `yaml:"target_audience"` // Целевая аудитория
		Status         Status          `yaml:"status"`          // Статус агрегата
		Outmembers     []Outmember     `yaml:"-"`               // Согласования
		Links          []Link          `yaml:"links"`           // Ссылки
		CreationDate   time.Time       `yaml:"creation_date"`   // 
		CreatedAt      time.Time       `yaml:"created_at"`      // 
		UpdatedAt      time.Time       `yaml:"updated_at"`      // 
		Bookings       []Booking       `yaml:"booking_ids"`     // 
		Slots        []Slot        `db:"-" yaml:"-"`        // Бюджеты
		Equipments        []Equipment        `db:"-" yaml:"-"`        // 
		Releases       []Release       `db:"-" yaml:"-"`        // Релизы
		Attachments    []Attachment    `yaml:"attachments"`     // Вложения
		Participants   []Participant   `yaml:"participants"`    // Участники
		Reports        []ReportRoom `yaml:"reports"`         // Отчеты
		State          State           `yaml:"state"`  
		Space          int32     `yaml:"space"`
	    SecurityEmail  string    `yaml:"security_email"`
	    Visible        bool    `yaml:"visible"`         
	}

	RoomModelList []Room
)

func (p Room) IsAssignee(user User) bool {
	return p.Employee != nil && p.Employee.ID == user.ID
}

func (p Room) IsOwner(user User) bool {
	return p.Owner != nil && p.Owner.ID == user.ID
}

func (p Room) IsAuthor(user User) bool {
	return p.Author != nil && p.Author.ID == user.ID
}

func (p Room) IsParticipant(user User) bool {
	if p.Participants != nil {
		for _, participant := range p.Participants {
			if participant.User.ID == user.ID {
				return true
			}
		}
	}

	return false
}
