package dbmodel

import (
	"time"

	"github.com/google/uuid"

	"be/internal/model"
)

type (
	Room struct {
		ID             uuid.UUID         `db:"id" yaml:"id"`
		Number         uint64            `db:"number" yaml:"number"`
		Owner          User              `db:"owner" yaml:"-"`
		Author         User              `db:"author" yaml:"-"`
		Employee       User              `db:"employee" yaml:"-"`
		Title          string            `db:"title" yaml:"title"`
		Description    string            `db:"description" yaml:"description"`
		TargetAudience string            `db:"target_audience" yaml:"target_audience"`
		Status         model.Status      `db:"status" yaml:"status"`
		Bookings       []uuid.UUID       `db:"booking_ids" yaml:"booking_ids"`
		BookingsModel  []model.Booking   `db:"-"`
		Outmembers     []model.Outmember `db:"-" yaml:"-"`
		Slots        SlotList        `db:"-" yaml:"-"`
		Equipments        EquipmentList        `db:"-" yaml:"-"`
		Releases       ReleaseList       `db:"-" yaml:"-"`
		Links          []model.Link      `db:"links"`
		CreationDate   time.Time         `db:"creation_date" yaml:"creation_date"`
		CreatedAt      time.Time         `db:"created_at" yaml:"created_at"`
		UpdatedAt      time.Time         `db:"updated_at" yaml:"updated_at"`
		Attachments    AttachmentList    `db:"-"`
		Reports        ReportRoomList `db:"-"`
		Participants   ParticipantList   `db:"participants"`
		State          model.State       `db:"state"`
		Space          int32     `db:"space"`
	    SecurityEmail  string    `db:"security_email"`
	    Visible        bool    `db:"visible"`
	}

	RoomList []Room
)

func RoomFromModel(p model.Room) Room {
	dbRoom := Room{
		ID:             p.ID,
		Number:         p.Number,
		Title:          p.Title,
		Description:    p.Description,
		Outmembers:     p.Outmembers,
		TargetAudience: p.TargetAudience,
		Status:         p.Status,
		Links:          p.Links,
		CreationDate:   p.CreationDate,
		Slots:        ToSlotList(p.Slots),
		Equipments:        ToEquipmentList(p.Equipments),
		Releases:       ToReleaseList(p.Releases),
		CreatedAt:      p.CreatedAt,
		UpdatedAt:      p.UpdatedAt,
		Attachments:    AttachmentListFromModel(p.Attachments),
		Reports:        ToReportRoomList(p.Reports),
		State:          p.State,
		Participants:   ParticipantsFromModel(p.Participants),
		Space:          p.Space,
	    SecurityEmail:  p.SecurityEmail,
	    Visible:        p.Visible,
	}

	if p.Owner != nil {
		dbRoom.Owner = UserFromModel(*p.Owner)
	}

	if p.Author != nil {
		dbRoom.Author = UserFromModel(*p.Author)
	}

	if p.Employee != nil {
		dbRoom.Employee = UserFromModel(*p.Employee)
	}

	return dbRoom
}

func (p Room) ToModel() model.Room {
	room := model.Room{
		ID:             p.ID,
		Number:         p.Number,
		Title:          p.Title,
		Description:    p.Description,
		TargetAudience: p.TargetAudience,
		Status:         p.Status,
		Outmembers:     p.Outmembers,
		Links:          p.Links,
		CreationDate:   p.CreationDate,
		CreatedAt:      p.CreatedAt,
		UpdatedAt:      p.UpdatedAt,
		Bookings:       p.BookingsModel,
		Slots:        p.Slots.Slots(),
		Equipments:        p.Equipments.Equipments(),
		Releases:       p.Releases.Releases(),
		Attachments:    p.Attachments.Attachments(),
		Reports:        p.Reports.ReportsRoom(),
		State:          p.State,
		Space:           p.Space,
	    SecurityEmail:   p.SecurityEmail,
	    Visible:         p.Visible,
	}

	if p.Owner.ID != nil {
		tmp := p.Owner.ToModel()
		room.Owner = &tmp
	}

	if p.Author.ID != nil {
		tmp := p.Author.ToModel()
		room.Author = &tmp
	}

	if p.Employee.ID != nil {
		tmp := p.Employee.ToModel()
		room.Employee = &tmp
	}

	return room
}

func (p Room) ToModelPtr() *model.Room {
	pPtr := p.ToModel()
	return &pPtr
}

func (l RoomList) Rooms() []model.Room {
	modelsList := make([]model.Room, 0, len(l))
	for _, room := range l {
		modelsList = append(modelsList, room.ToModel())
	}

	return modelsList
}

func (p Room) ParticipantUUIDs() []uuid.UUID {
	participantUUIDs := make([]uuid.UUID, 0, len(p.Participants))
	for _, participant := range p.Participants {
		participantUUIDs = append(participantUUIDs, participant.UserID)
	}

	return participantUUIDs
}
