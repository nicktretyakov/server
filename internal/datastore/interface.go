package datastore

import (
	"context"

	"github.com/google/uuid"

	"be/internal/datastore/filters"
	addressStore "be/internal/datastore/address"
	"be/internal/datastore/sorting"
	"be/internal/model"
)

type IDatastore interface {
	User() IUserStore
	Session() ISessionStore
	Booking() IBookingStore
	Report() IReportStore
	FinalReport() IFinalReportStore
	Outmember() IOutmemberStore
	Dictionary() IDictionaryStore
	Stage() IStageStore
	Address() IAddressStore
	Room() IRoomStore
}

type IUserStore interface {
	CreateUser(ctx context.Context, user model.User) (*model.User, error)
	UpdateUser(ctx context.Context, user model.User) (*model.User, error)
	FindUserByProfileID(ctx context.Context, uid string) (*model.User, error)
	FindUserByPortalCode(ctx context.Context, uid uint64) (*model.User, error)
	FindUsersIDByPortalCode(ctx context.Context, uid []uint64) ([]uuid.UUID, error)
	FindUserByPK(ctx context.Context, id uuid.UUID) (*model.User, error)
	FindUsersByIDs(ctx context.Context, id []uuid.UUID) ([]model.User, error)
	FindAdmins(ctx context.Context) ([]*model.User, error)

	CreateNoteSettings(ctx context.Context, userID *uuid.UUID, emailOn, lifeOn bool) (*model.NoteSettings, error)
	GetNoteSettingByUser(ctx context.Context, userID *uuid.UUID) (*model.NoteSettings, error)
	SetEmailNoteSetting(ctx context.Context, userID *uuid.UUID, emailOn bool) error
	SetLifeNoteSetting(ctx context.Context, userID *uuid.UUID, lifeOn bool) error
}

type ISessionStore interface {
	FindSessionByRefreshToken(ctx context.Context, refresh string) (*model.Session, error)
	FindSessionByID(ctx context.Context, sessionID uuid.UUID) (*model.Session, error)
	RefreshSession(ctx context.Context, session model.Session) error
	NewSession(ctx context.Context, session model.Session) (*model.Session, error)
}

type IBookingStore interface {
	FindByID(ctx context.Context, bookingID uuid.UUID) (*model.Booking, error)
	Create(ctx context.Context, booking model.Booking) (*model.Booking, error)
	Update(ctx context.Context, booking model.Booking) (*model.Booking, error)
	CreateBookingUser(ctx context.Context, bookingUser model.BookingUser) error
	UpdateBookingUser(ctx context.Context, bookingUserBase, bookingUserChange model.BookingUser) error
	DeleteBookingUser(ctx context.Context, bookingUser model.BookingUser) error
	GetListSupervisors(ctx context.Context) ([]uuid.UUID, error)
	UpdateStatus(ctx context.Context, bookingID uuid.UUID, status model.Status) error
	UpdateLinks(ctx context.Context, bookingID uuid.UUID, link []model.Link) error
	UpdateState(ctx context.Context, bookingID uuid.UUID, state model.State) error
	ActiveList(ctx context.Context, limit, offset uint64,
		sorting sorting.Sorting, queryFilters ...filters.Filter) ([]model.Booking, uint64, error)
	ListBookingsWithNotSendReports(ctx context.Context) ([]model.Booking, error)
}

type IRoomStore interface {
	FindByID(ctx context.Context, roomID uuid.UUID) (*model.Room, error)
	FindSlotByID(ctx context.Context, slotID uuid.UUID) (*model.Slot, error)
	FindEquipmentByID(ctx context.Context, equipmentID uuid.UUID) (*model.Equipment, error)
	FindReleaseByID(ctx context.Context, release uuid.UUID) (*model.Release, error)
	FindSlotsByID(ctx context.Context, slotID []uuid.UUID) ([]model.Slot, error)
	FindEquipmentsByID(ctx context.Context, equipmentID []uuid.UUID) ([]model.Equipment, error)
	FindReleasesByID(ctx context.Context, releaseID []uuid.UUID) ([]model.Release, error)
	FindRoomObjectsByTimeline(ctx context.Context, roomID uuid.UUID,
		period model.Timeline) ([]model.Slot, []model.Equipment, []model.Release, error)
	Create(ctx context.Context, room model.Room, equipmentIDs, slotIDs []uuid.UUID) (*model.Room, error)
	Update(ctx context.Context, room model.Room, equipmentIDs, slotIDs, bookingIDs []uuid.UUID) (*model.Room, error)
	AddSlots(ctx context.Context, slots []model.Slot) ([]model.Slot, error)
	UpdateSlot(ctx context.Context, slot model.Slot) (*model.Slot, error)
	DeleteSlot(ctx context.Context, slotID uuid.UUID) error
	AddEquipments(ctx context.Context, equipments []model.Equipment) ([]model.Equipment, error)
	UpdateEquipment(ctx context.Context, equipment model.Equipment) (*model.Equipment, error)
	DeleteEquipment(ctx context.Context, equipmentID uuid.UUID) error
	AddReleases(ctx context.Context, releases []model.Release) ([]model.Release, error)
	UpdateRelease(ctx context.Context, release model.Release) (*model.Release, error)
	DeleteRelease(ctx context.Context, releaseID uuid.UUID) error
	UpdateBookingsRoom(ctx context.Context, roomID uuid.UUID, bookingIDs []uuid.UUID) error
	ActiveList(ctx context.Context, limit, offset uint64,
		sorting sorting.Sorting, queryFilters ...filters.Filter) ([]model.Room, uint64, error)
	UpdateLinks(ctx context.Context, roomID uuid.UUID, link []model.Link) error
	UpdateStatus(ctx context.Context, roomID uuid.UUID, status model.Status) error
	UpdateParticipant(ctx context.Context, roomID uuid.UUID, participants map[uuid.UUID]string) error
	UpdateState(ctx context.Context, roomID uuid.UUID, state model.State) error
}

type IAddressStore interface {
	CreateAttachment(ctx context.Context, attachment model.Attachment) (*model.Attachment, error)
	FindAttachmentByID(ctx context.Context, attachmentID uuid.UUID) (*model.Attachment, error)
	DeleteAttachment(ctx context.Context, attachment model.Attachment) error
	RenameAttachment(ctx context.Context, attachment model.Attachment) (*model.Attachment, error)

	CreateEmailNote(ctx context.Context, note model.EmailNote) (*model.EmailNote, error)
	UpdateEmailNote(ctx context.Context, note model.EmailNote) (*model.EmailNote, error)
	EmailNotesList(ctx context.Context) ([]*model.EmailNote, error)

	LifeNotesList(ctx context.Context) ([]*model.LifeNote, error)
	CreateLifeNote(ctx context.Context, note model.LifeNote) (*model.LifeNote, error)
	UpdateLifeNote(ctx context.Context, note model.LifeNote) (*model.LifeNote, error)
	GetChannelFromLifeNote(ctx context.Context, userID *uuid.UUID) (uuid.UUID, error)

	SystemNoteList(ctx context.Context, userID uuid.UUID,
		status model.SystemStatus, limit, offset uint32, sorting sorting.Sorting) ([]*model.SystemNote, error)
	GetSystemNotesCount(ctx context.Context, userID uuid.UUID) (int, error)
	UpdateSystemNotes(ctx context.Context, noteIDs []uuid.UUID) error
	CreateSystemNote(ctx context.Context, note model.SystemNote) (*model.SystemNote, error)

	UploadFileExportedAddresss(ctx context.Context, attachment model.Attachment) (*model.Attachment, error)
}

type IReportStore interface {
	FindBookingReportByID(ctx context.Context, report uuid.UUID) (*model.ReportBooking, error)
	Update(ctx context.Context, report model.ReportBooking) (*model.ReportBooking, error)
	BulkCreate(ctx context.Context, reports []model.ReportBooking) ([]model.ReportBooking, error)
	CreateReportRoom(ctx context.Context, report model.ReportRoom) (*model.ReportRoom, error)
	ListByBookingID(ctx context.Context, bookingID uuid.UUID) ([]model.ReportBooking, error)
	BulkDelete(ctx context.Context, reports []model.ReportBooking) error
	FindRoomReportByID(ctx context.Context, report uuid.UUID) (*model.ReportRoom, error)
}

type IFinalReportStore interface {
	FindByID(ctx context.Context, reportID uuid.UUID) (*model.FinalReport, error)
	FindByBookingID(ctx context.Context, bookingID uuid.UUID) (*model.FinalReport, error)
	Update(ctx context.Context, report model.FinalReport) (*model.FinalReport, error)
	UpdateStatus(ctx context.Context, bookingID uuid.UUID, status model.FinalReportStatus) error
	Create(ctx context.Context, report model.FinalReport) (*model.FinalReport, error)
}

type IDictionaryStore interface {
	DepartmentList(ctx context.Context) ([]model.Department, error)
}

type IOutmemberStore interface {
	Create(ctx context.Context, outmember model.Outmember, typeAddress addressStore.TypeAddress) (*model.Outmember, error)
	BulkDelete(ctx context.Context, addressID uuid.UUID, agrType model.OutmemberType,
		typeAddress addressStore.TypeAddress) error
	List(ctx context.Context, typeAddress addressStore.TypeAddress, filters ...filters.Filter) ([]model.Outmember, error)
}

type IStageStore interface {
	FindByID(ctx context.Context, stage uuid.UUID) (*model.Stage, error)
	ListByBookingID(ctx context.Context, bookingID uuid.UUID) ([]model.Stage, error)
	Create(ctx context.Context, stage model.Stage) (*model.Stage, error)
	Update(ctx context.Context, stage model.Stage) (*model.Stage, error)
	Remove(ctx context.Context, stage model.Stage) error

	FindStageIssueByID(ctx context.Context, issueID uuid.UUID) (*model.Issue, error)
	CreateIssue(ctx context.Context, issue model.Issue) (*model.Issue, error)
	UpdateIssue(ctx context.Context, issue model.Issue) error
	RemoveIssue(ctx context.Context, issue model.Issue) error
}

type IIssueStore interface {
	FindByID(ctx context.Context, issueID uuid.UUID) (*model.Issue, error)
	Create(ctx context.Context, issue model.Issue) (*model.Issue, error)
	Update(ctx context.Context, issue model.Issue) error
	Remove(ctx context.Context, issueID uuid.UUID) error
}
