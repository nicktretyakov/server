package model

import (
	"time"

	"github.com/google/uuid"
)

type BookingType int

const (
	InvestBookingType         BookingType = 1
	OrganizationalBookingType BookingType = 2

	ColumnNumberBooking      = "Номер проекта"
	ColumnBookingTitle       = "Название проекта"
	ColumnBookingDepartments = "Подразделение"
	ColumnBookingCity        = "Город"
	ColumnBookingSlot      = "Бюджет проекта (руб.)"
	ColumnBookingStartAt     = "Начало проекта"
	ColumnBookingEndAt       = "Первоначальный срок окончания"
	ColumnBookingSupervisor  = "Руководитель проекта"
	ColumnBookingStatus      = "Текущий статус проекта"

	BookingColumnLowest         = "A"
	BookingColumnHighest        = "I"
	BookingPrefixOutputFilename = "Bookings"
)

func (t BookingType) IsValid() bool {
	return t == InvestBookingType || t == OrganizationalBookingType
}

type (
	Booking struct {
		ID           uuid.UUID       `yaml:"id"`            //  ID проекта
		Number       uint64          `yaml:"number"`        //  Номер проекта
		CreatedAt    time.Time       `yaml:"created_at"`    //  Ts создания
		UpdatedAt    time.Time       `yaml:"updated_at"`    //  Ts обновления
		Type         BookingType     `yaml:"type"`          //  Тип проекта
		Title        string          `yaml:"title"`         //  Имя проекта
		City         string          `yaml:"city"`          //  Город
		Timeline     Timeline        `yaml:"timeline"`      //  Срок проекта
		Description  string          `yaml:"description"`   //  Описание проекта
		Tags         []Tag           `yaml:"tags"`          //  Теги
		Attachments  []Attachment    `yaml:"attachments"`   //  Вложения
		Departments  []Department    `yaml:"departments"`   //  Подразделения
		Reports      []ReportBooking `yaml:"reports"`       //  Отчеты
		FinalReport  FinalReport     `yaml:"-"`             //  Финальный отчет
		Slot       Notification           `yaml:"slot"`        //  Бюджет проекта
		Supervisor   *User           `yaml:"supervisor"`    //  Руководитель проекта
		Author       *User           `yaml:"author"`        //  Автор проекта
		Assignee     *User           `yaml:"room_owner"` //  Согласующий проекта
		Participants []Participant   `yaml:"participants"`  //  Участники проекта
		Goal         string          `yaml:"goal"`          //  Цель проекта
		Stages       []Stage         `yaml:"stages"`        //  Этапы проекта
		Status       Status          `yaml:"status"`        //  Статус агрегата
		Outmembers   []Outmember     `yaml:"-"`
		Links        []Link          `yaml:"links"` // Ссылки
		State        State           `yaml:"state"` // Состояние
	}

	BookingsModelList []Booking
)

func (p Booking) IsSupervisor(user User) bool {
	return p.Supervisor != nil && p.Supervisor.ID == user.ID
}

func (p Booking) IsAssignee(user User) bool {
	return p.Assignee != nil && p.Assignee.ID == user.ID
}

func (p Booking) IsParticipant(user User) bool {
	if p.Participants != nil {
		for _, participant := range p.Participants {
			if participant.User.ID == user.ID {
				return true
			}
		}
	}

	return false
}

func (p Booking) IsAuthor(user User) bool {
	return p.Author != nil && p.Author.ID == user.ID
}

func (p Booking) IsHeadOfBooking(user User) bool {
	return user.Role == Admin
}

func (p Booking) IsCEO(user User) bool {
	return user.Role == CEO
}

func (p Booking) IsPartner(user User) bool {
	return user.Role == Partner
}

func (p Booking) IsAssigned() bool {
	return p.Assignee != nil
}

func (p Booking) IsFinalReport() bool {
	return p.FinalReport.ID != uuid.Nil
}

func (p Booking) GetAttachmentsFinalReport() (attachmentsFinalReport []Attachment) {
	attachmentsFinalReport = make([]Attachment, 0, len(p.Attachments))

	if p.FinalReport.ID == uuid.Nil {
		return
	}

	for _, prAttachment := range p.Attachments {
		for _, finalReportAttachmentUUID := range p.FinalReport.AttachmentsUUID {
			if finalReportAttachmentUUID == prAttachment.ID {
				attachmentsFinalReport = append(attachmentsFinalReport, prAttachment)

				break
			}
		}
	}

	return
}
