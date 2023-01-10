package testutil

import (
	"fmt"
	"path"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"be/internal/datastore/base"
	"be/internal/datastore/dbmodel"
	"be/internal/model"
)

var ExistedUserID = uuid.MustParse("5ba5b3d2-87c4-4a72-90dd-cbe192e00150") //nolint:gochecknoglobals

//nolint:funlen
func NewFixtureStore(t *testing.T) *FixtureStore {
	t.Helper()

	source := "../testutil/fixtures"

	yamlBookings := make(bookingList, 0)
	require.NoError(t, readFile(path.Join(source, "bookings.yml"), &yamlBookings))

	dbUsers := make(dbmodel.UserList, 0)
	require.NoError(t, readFile(path.Join(source, "users.yml"), &dbUsers))

	bookingUsers := make([]model.BookingUser, 0)
	require.NoError(t, readFile(path.Join(source, "booking_users.yml"), &bookingUsers))

	departments := make([]model.Department, 0)
	require.NoError(t, readFile(path.Join(source, "dictionary_departments.yml"), &departments))

	bookingDepartments := make([]bookingDepartment, 0)
	require.NoError(t, readFile(path.Join(source, "booking_departments.yml"), &bookingDepartments))

	bookingAttachments := make(dbmodel.AttachmentList, 0)
	require.NoError(t, readFile(path.Join(source, "attachments.yml"), &bookingAttachments))

	bookingOutmembers := make([]model.Outmember, 0)
	require.NoError(t, readFile(path.Join(source, "booking_outmember.yml"), &bookingOutmembers))

	bookingFinalReports := make([]model.FinalReport, 0)
	require.NoError(t, readFile(path.Join(source, "booking_final_reports.yml"), &bookingFinalReports))

	bookingStages := make([]dbmodel.Stage, 0)
	require.NoError(t, readFile(path.Join(source, "booking_stages.yml"), &bookingStages))

	stageIssues := make([]dbmodel.Issue, 0)
	require.NoError(t, readFile(path.Join(source, "stage_issues.yml"), &stageIssues))

	bookingReports := make([]dbmodel.ReportBooking, 0)
	require.NoError(t, readFile(path.Join(source, "booking_reports.yml"), &bookingReports))

	yamlRooms := make(roomList, 0)
	require.NoError(t, readFile(path.Join(source, "rooms.yml"), &yamlRooms))

	roomSlot := make([]model.Slot, 0)
	require.NoError(t, readFile(path.Join(source, "room_slot.yml"), &roomSlot))

	roomEquipment := make([]model.Equipment, 0)
	require.NoError(t, readFile(path.Join(source, "room_equipment.yml"), &roomEquipment))

	roomRelease := make([]model.Release, 0)
	require.NoError(t, readFile(path.Join(source, "room_release.yml"), &roomRelease))

	roomReports := make([]model.ReportRoom, 0)
	require.NoError(t, readFile(path.Join(source, "room_reports.yml"), &roomReports))

	bookings := yamlBookings.Bookings()
	rooms := yamlRooms.Rooms()
	users := dbUsers.Users()

	setUserRoles(bookings, users, bookingUsers)
	setBookingDepartments(bookings, departments, bookingDepartments)
	setBookingReports(bookings, bookingReports)
	setBookingOutmembers(bookings, bookingOutmembers)
	setBookingFinalReports(bookings, bookingFinalReports)
	setBookingStages(bookings, bookingStages, stageIssues)
	setRoomSlot(rooms, roomSlot)
	setRoomEquipment(rooms, roomEquipment)
	setRoomReports(rooms, roomReports)
	setRoomReleases(rooms, roomRelease)

	return &FixtureStore{
		users:              users,
		bookings:           bookings,
		rooms:              rooms,
		bookingUsers:       bookingUsers,
		departments:        departments,
		bookingAttachments: bookingAttachments.Attachments(),
		slots:              roomSlot,
		equipments:         roomEquipment,
		roomReports:        roomReports,
		releases:           roomRelease,
	}
}

func setBookingDepartments(bookings []model.Booking, depsList []model.Department, bookDeps []bookingDepartment) {
	for _, bookDepartment := range bookDeps {
		dep := getDepartmentByID(depsList, bookDepartment.DepartmentID)
		_, i := getBookingByID(bookings, bookDepartment.BookingID)

		if len(bookings[i].Departments) == 0 {
			bookings[i].Departments = make([]model.Department, 0, 1)
		}

		bookings[i].Departments = append(bookings[i].Departments, dep)
	}
}

func setBookingReports(bookings []model.Booking, reports []dbmodel.ReportBooking) {
	for _, report := range reports {
		_, i := getBookingByID(bookings, report.BookingID)

		bookings[i].Reports = append(bookings[i].Reports, report.ToModel())
	}
}

func setBookingOutmembers(bookings []model.Booking, outmembers []model.Outmember) {
	for _, agr := range outmembers {
		_, i := getBookingByID(bookings, agr.AddressID)

		if len(bookings[i].Outmembers) == 0 {
			bookings[i].Outmembers = make([]model.Outmember, 0, 1)
		}

		bookings[i].Outmembers = append(bookings[i].Outmembers, agr)
	}
}

func setBookingFinalReports(bookings []model.Booking, finalReports []model.FinalReport) {
	for _, finalRep := range finalReports {
		_, i := getBookingByID(bookings, finalRep.BookingID)

		bookings[i].FinalReport = finalRep
	}
}

func setUserRoles(bookings []model.Booking, users []model.User, bookingUsers []model.BookingUser) {
	for i := range bookings {
		for _, user := range users {
			for _, bookingUser := range bookingUsers {
				setBookingUserRole(&bookings[i], user, bookingUser)
			}
		}
	}
}

func setRoomSlot(rooms []model.Room, slot []model.Slot) {
	for _, bud := range slot {
		_, i := getRoomByID(rooms, bud.RoomID)

		if len(rooms[i].Slots) == 0 {
			rooms[i].Slots = make([]model.Slot, 0, 1)
		}

		rooms[i].Slots = append(rooms[i].Slots, bud)
	}
}

func setRoomEquipment(rooms []model.Room, equipment []model.Equipment) {
	for _, met := range equipment {
		_, i := getRoomByID(rooms, met.RoomID)

		if len(rooms[i].Equipments) == 0 {
			rooms[i].Equipments = make([]model.Equipment, 0, 1)
		}

		rooms[i].Equipments = append(rooms[i].Equipments, met)
	}
}

func setRoomReports(rooms []model.Room, reports []model.ReportRoom) {
	for _, rep := range reports {
		_, i := getRoomByID(rooms, rep.RoomID)

		if len(rooms[i].Reports) == 0 {
			rooms[i].Reports = make([]model.ReportRoom, 0, 1)
		}

		rooms[i].Reports = append(rooms[i].Reports, rep)
	}
}

func setRoomReleases(rooms []model.Room, releases []model.Release) {
	for _, rel := range releases {
		_, i := getRoomByID(rooms, rel.RoomID)

		if len(rooms[i].Releases) == 0 {
			rooms[i].Releases = make([]model.Release, 0, 1)
		}

		rooms[i].Releases = append(rooms[i].Releases, rel)
	}
}

func setBookingUserRole(booking *model.Booking, user model.User, bookingUser model.BookingUser) {
	if bookingUser.BookingID == booking.ID &&
		user.ID == bookingUser.UserID {
		u := user

		switch bookingUser.Role {
		case model.SupervisorBookingUserRole:
			booking.Supervisor = &u
		case model.AssigneeBookingUserRole:
			booking.Assignee = &u
		case model.AuthorBookingUserRole:
			booking.Author = &u
		case model.ParticipantBookingUserRole:
			booking.Participants = append(booking.Participants, model.Participant{
				User: user,
				Role: "role title",
			})
		}
	}
}

func setBookingStages(bookings []model.Booking, stages []dbmodel.Stage, issues []dbmodel.Issue) {
	for _, stageRep := range stages {
		if stageRep.Status != model.ActiveAggregateStatus {
			continue
		}

		_, ipr := getBookingByID(bookings, stageRep.BookingID)

		bookings[ipr].Stages = append(bookings[ipr].Stages, stageRep.ToModel())

		for _, issueRep := range issues {
			if issueRep.StageID != stageRep.ID {
				continue
			}

			_, i := getStageByID(bookings[ipr].Stages, issueRep.StageID)

			bookings[ipr].Stages[i].Issues = append(bookings[ipr].Stages[i].Issues, issueRep.ToModel())
		}
	}
}

type FixtureStore struct {
	users              []model.User
	bookings           []model.Booking
	bookingUsers       []model.BookingUser
	departments        []model.Department
	bookingAttachments []model.Attachment
	rooms              []model.Room
	slots              []model.Slot
	equipments         []model.Equipment
	roomReports        []model.ReportRoom
	releases           []model.Release
}

func (s FixtureStore) GetBookings() []model.Booking {
	return s.bookings
}

func (s FixtureStore) PickBookings(idx ...int) []model.Booking {
	w := make([]model.Booking, 0, len(idx))
	for _, i := range idx {
		w = append(w, s.bookings[i])
	}

	return w
}

func (s FixtureStore) PickRooms(idx ...int) []model.Room {
	w := make([]model.Room, 0, len(idx))
	for _, i := range idx {
		w = append(w, s.rooms[i])
	}

	return w
}

func (s FixtureStore) GetBookingByID(bookingID uuid.UUID) model.Booking {
	p, _ := getBookingByID(s.bookings, bookingID)

	return p
}

func (s FixtureStore) GetRoomByID(roomID uuid.UUID) model.Room {
	room, _ := getRoomByID(s.rooms, roomID)

	return room
}

func (s FixtureStore) GetSlotByID(slotID uuid.UUID) model.Slot {
	for _, bud := range s.slots {
		if bud.ID == slotID {
			return bud
		}
	}

	panic(fmt.Sprintf("slot %s is not in fixtures", slotID.String()))
}

func (s FixtureStore) GetEquipmentByID(equipmentID uuid.UUID) model.Equipment {
	for _, met := range s.equipments {
		if met.ID == equipmentID {
			return met
		}
	}

	panic(fmt.Sprintf("equipment %s is not in fixtures", equipmentID.String()))
}

func (s FixtureStore) GetReleaseByID(releaseID uuid.UUID) model.Release {
	for _, rel := range s.releases {
		if rel.ID == releaseID {
			return rel
		}
	}

	panic(fmt.Sprintf("equipment %s is not in fixtures", releaseID.String()))
}

func (s FixtureStore) GetRoomReportByID(reportID uuid.UUID) model.ReportRoom {
	for _, rep := range s.roomReports {
		if rep.ID == reportID {
			return rep
		}
	}

	panic(fmt.Sprintf("room report %s is not in fixtures", reportID.String()))
}

func getBookingByID(bookings []model.Booking, bookingID uuid.UUID) (model.Booking, int) {
	for i, book := range bookings {
		if book.ID == bookingID {
			return book, i
		}
	}

	panic(fmt.Sprintf("booking %s is not in fixtures", bookingID.String()))
}

func getStageByID(stages []model.Stage, stageID uuid.UUID) (model.Stage, int) {
	for i, st := range stages {
		if st.ID == stageID {
			return st, i
		}
	}

	panic(fmt.Sprintf("stage %s is not in fixtures", stageID.String()))
}

func getRoomByID(rooms []model.Room, roomID uuid.UUID) (model.Room, int) {
	for i, room := range rooms {
		if room.ID == roomID {
			return room, i
		}
	}

	panic(fmt.Sprintf("room %s is not in fixtures", roomID.String()))
}

func (s FixtureStore) GetDepartmentByID(id uuid.UUID) model.Department {
	return getDepartmentByID(s.departments, id)
}

func getDepartmentByID(departments []model.Department, id uuid.UUID) model.Department {
	for _, dep := range departments {
		if dep.ID == id {
			return dep
		}
	}

	panic("department is not fixtures")
}

func (s FixtureStore) GetStageByID(stageID uuid.UUID) model.Stage {
	for _, booking := range s.bookings {
		for _, stage := range booking.Stages {
			if stage.ID == stageID {
				return stage
			}
		}
	}

	panic("stage is not fixtures")
}

func (s FixtureStore) GetIssueByID(issueID uuid.UUID) model.Issue {
	for _, booking := range s.bookings {
		for _, stage := range booking.Stages {
			for _, issue := range stage.Issues {
				if issue.ID == issueID {
					return issue
				}
			}
		}
	}

	panic("issue is not fixtures")
}

func (s FixtureStore) GetUsers() []model.User {
	return s.users
}

func (s FixtureStore) UserByID(wantID uuid.UUID) (model.User, error) {
	for _, user := range s.users {
		if user.ID == wantID {
			return user, nil
		}
	}

	return model.User{}, base.ErrNotFound
}

func (s FixtureStore) MustUserByID(wantID uuid.UUID) model.User {
	u, err := s.UserByID(wantID)
	if err != nil {
		panic(err)
	}

	return u
}

func (s FixtureStore) MustAttachmentByID(wantID uuid.UUID) model.Attachment {
	a, err := s.attachmentByID(wantID)
	if err != nil {
		panic(err)
	}

	return a
}

func (s FixtureStore) GetReportID() uuid.UUID {
	return uuid.MustParse("15cb22fa-b363-4235-899c-282277470f3a")
}

func (s FixtureStore) attachmentByID(id uuid.UUID) (model.Attachment, error) {
	for _, attachment := range s.bookingAttachments {
		if attachment.ID == id {
			return attachment, nil
		}
	}

	return model.Attachment{}, base.ErrNotFound
}
