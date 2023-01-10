package room

import (
	"context"
	"errors"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"be/internal/datastore/base"
	"be/internal/datastore/dbmodel"
	"be/internal/datastore/filters"
	"be/internal/datastore/outmember"
	"be/internal/datastore/sorting"
	"be/internal/model"
)

func (s *Storage) ActiveList(
	ctx context.Context,
	limit, offset uint64,
	sorting sorting.Sorting,
	queryFilters ...filters.Filter,
) ([]model.Room, uint64, error) {
	queryFilters = append(queryFilters, newStatusNegFilter([]model.Status{model.InitialAddressStatus}))
	query, countQuery := withFilters(limit, offset, queryFilters...)
	query = sorting.Apply(query)

	var (
		rooms = make(dbmodel.RoomList, 0)
		count uint64
	)

	if err := s.db.Select(ctx, query, &rooms); err != nil {
		return nil, 0, err
	}

	if err := s.db.Get(ctx, countQuery, &count); err != nil {
		return nil, 0, err
	}

	return rooms.Rooms(), count, nil
}

func withFilters(limit, offset uint64, queryFilters ...filters.Filter) (sq.SelectBuilder, sq.SelectBuilder) {
	query := roomSelectQuery()
	countQuery := roomCountQuery()

	for _, filter := range queryFilters {
		query = filter.Apply(query)
		countQuery = filter.Apply(countQuery)
	}

	if limit > 0 {
		query = newLimitFilter(limit).Apply(query)
	}

	if offset > 0 {
		query = newOffsetFilter(offset).Apply(query)
	}

	return query, countQuery
}

//nolint:gocognit,gocyclo,cyclop,funlen
func (s *Storage) FindByID(ctx context.Context, roomID uuid.UUID) (*model.Room, error) {
	query := roomSelectQuery().Where("p.id=?", roomID)

	var (
		room           dbmodel.Room
		roomOutmembers []model.Outmember
		roomSlots      dbmodel.SlotList
		roomEquipments dbmodel.EquipmentList
		bookings       dbmodel.BookingList
		roomReleases   dbmodel.ReleaseList
		roomReports    dbmodel.ReportRoomList
	)

	if err := s.db.Get(ctx, query, &room); err != nil {
		return nil, err
	}

	if err := s.db.Select(ctx, roomOutmembersSelectQuery(roomID), &roomOutmembers); err != nil && !errors.Is(err, base.ErrNotFound) {
		return nil, err
	}

	if err := s.db.Select(ctx, roomSlotsSelectQuery(roomID), &roomSlots); err != nil && !errors.Is(err, base.ErrNotFound) {
		return nil, err
	}

	if err := s.db.Select(ctx, roomEquipmentsSelectQuery(roomID), &roomEquipments); err != nil && !errors.Is(err, base.ErrNotFound) {
		return nil, err
	}

	if err := s.db.Select(ctx, roomReleaseSelectQuery(roomID), &roomReleases); err != nil && !errors.Is(err, base.ErrNotFound) {
		return nil, err
	}

	if err := s.db.Select(ctx, bookingWithUsersSelectQuery(room.Bookings), &bookings); err != nil && !errors.Is(err, base.ErrNotFound) {
		return nil, err
	}

	if err := s.db.Select(ctx, reportsSelectQuery(roomID), &roomReports); err != nil && !errors.Is(err, base.ErrNotFound) {
		return nil, err
	}

	roomAttachments, err := s.getAttachmentsModelDB(ctx, roomID)
	if err != nil {
		return nil, err
	}

	room.Outmembers = roomOutmembers
	room.Slots = roomSlots
	room.Equipments = roomEquipments
	room.Attachments = roomAttachments
	room.BookingsModel = bookings.Bookings()
	room.Releases = roomReleases
	room.Reports = roomReports

	roomModel := room.ToModel()

	participants, err := s.setParticipants(ctx, room)
	if err != nil {
		return nil, err
	}

	roomModel.Participants = participants

	return &roomModel, nil
}

func (s *Storage) setParticipants(ctx context.Context, room dbmodel.Room) ([]model.Participant, error) {
	var participants dbmodel.ParticipantList

	if err := s.db.Select(
		ctx,
		roomParticipantsSelectQuery(room.ID),
		&participants,
	); err != nil && !errors.Is(err, base.ErrNotFound) {
		return nil, err
	}

	return participants.Participants(), nil
}

func (s *Storage) FindRoomObjectsByTimeline(
	ctx context.Context,
	roomID uuid.UUID,
	period model.Timeline,
) ([]model.Slot, []model.Equipment, []model.Release, error) {
	var (
		roomSlots      dbmodel.SlotList
		roomEquipments dbmodel.EquipmentList
		roomReleases   dbmodel.ReleaseList
	)

	if err := s.db.Select(
		ctx,
		slotsByTimelineSelectQuery(roomID, period),
		&roomSlots,
	); err != nil && !errors.Is(err, base.ErrNotFound) {
		return nil, nil, nil, err
	}

	if err := s.db.Select(
		ctx,
		equipmentsTimelineSelectQuery(roomID, period),
		&roomEquipments,
	); err != nil && !errors.Is(err, base.ErrNotFound) {
		return nil, nil, nil, err
	}

	if err := s.db.Select(
		ctx,
		releasesTimelineSelectQuery(roomID, period),
		&roomReleases,
	); err != nil && !errors.Is(err, base.ErrNotFound) {
		return nil, nil, nil, err
	}

	return roomSlots.Slots(), roomEquipments.Equipments(), roomReleases.Releases(), nil
}

func (s *Storage) FindSlotByID(ctx context.Context, slotID uuid.UUID) (*model.Slot, error) {
	var slot dbmodel.Slot

	if err := s.db.Get(ctx, slotSelectQuery(slotID), &slot); err != nil {
		return nil, err
	}

	slotModel := slot.ToModel()

	return &slotModel, nil
}

func (s *Storage) FindEquipmentByID(ctx context.Context, equipmentID uuid.UUID) (*model.Equipment, error) {
	var equipment dbmodel.Equipment

	if err := s.db.Get(ctx, equipmentSelectQuery(equipmentID), &equipment); err != nil {
		return nil, err
	}

	equipmentModel := equipment.ToModel()

	return &equipmentModel, nil
}

func (s *Storage) FindReleaseByID(ctx context.Context, releaseID uuid.UUID) (*model.Release, error) {
	var release dbmodel.Release

	if err := s.db.Get(ctx, releaseSelectQuery(releaseID), &release); err != nil {
		return nil, err
	}

	releaseModel := release.ToModel()

	return &releaseModel, nil
}

func (s *Storage) FindSlotsByID(ctx context.Context, slotID []uuid.UUID) ([]model.Slot, error) {
	var roomSlots dbmodel.SlotList

	if err := s.db.Select(ctx, slotsSelectQuery(slotID), &roomSlots); err != nil && !errors.Is(err, base.ErrNotFound) {
		return nil, err
	}

	return roomSlots.Slots(), nil
}

func (s *Storage) FindEquipmentsByID(ctx context.Context, equipmentID []uuid.UUID) ([]model.Equipment, error) {
	var roomEquipments dbmodel.EquipmentList

	if err := s.db.Select(ctx, equipmentsSelectQuery(equipmentID), &roomEquipments); err != nil && !errors.Is(err, base.ErrNotFound) {
		return nil, err
	}

	return roomEquipments.Equipments(), nil
}

func (s *Storage) FindReleasesByID(ctx context.Context, releaseID []uuid.UUID) ([]model.Release, error) {
	var roomRelease dbmodel.ReleaseList

	if err := s.db.Select(ctx, releasesSelectQuery(releaseID), &roomRelease); err != nil && !errors.Is(err, base.ErrNotFound) {
		return nil, err
	}

	return roomRelease.Releases(), nil
}

func (s *Storage) getAttachmentsModelDB(ctx context.Context, addressID uuid.UUID) ([]dbmodel.Attachment, error) {
	var invObjAttachments []dbmodel.Attachment
	if err := s.db.Select(ctx, attachmentsSelectQuery().Where("address_id=?", addressID), &invObjAttachments); err != nil {
		return nil, err
	}

	for i, att := range invObjAttachments {
		uri, err := s.fileStorage.Link(att.Key, s.linkTimeLife)
		if err != nil {
			return nil, err
		}

		invObjAttachments[i].URL = uri
	}

	return invObjAttachments, nil
}

func attachmentsSelectQuery() sq.SelectBuilder {
	return base.Builder().
		Select(
			"attachments.id",
			"attachments.author_id",
			"attachments.created_at",
			"attachments.address_id",
			"attachments.key",
			"attachments.file_name",
			"attachments.mime",
			"attachments.size",
		).
		From(attachmentsTableName)
}

func roomSelectQuery() sq.SelectBuilder {
	selectColumns := []string{
		"p.id",
		"p.created_at",
		"p.updated_at",
		"p.number",
		"p.title",
		"p.description",
		"p.target_audience",
		"p.status",
		"p.links",
		"p.booking_ids",
		"p.creation_date",
		"p.participants",
		"p.state",
	}

	selectColumns = append(selectColumns, getUserColumns("u1", "author")...)
	selectColumns = append(selectColumns, getUserColumns("u2", "employee")...)
	selectColumns = append(selectColumns, getUserColumns("u3", "owner")...)

	return base.
		Builder().
		Select(selectColumns...).
		LeftJoin("users u1 on author_id = u1.id").
		LeftJoin("users u2 on employee_id = u2.id").
		LeftJoin("users u3 on owner_id = u3.id").
		From(roomTableNameAlias)
}

func roomCountQuery() sq.SelectBuilder {
	return base.
		Builder().
		Select("count(*)").
		From(roomTableNameAlias)
}

func roomOutmembersSelectQuery(roomID uuid.UUID) sq.SelectBuilder {
	selectColumns := []string{
		`room_outmember.id "id"`,
		`room_outmember.created_at "created_at"`,
		`room_outmember.type "type"`,
		`room_outmember.user_id "user_id"`,
		`room_outmember.address_id "address_id"`,
		`room_outmember.result "result"`,
		`room_outmember.extra "extra"`,
		`room_outmember.role "role"`,
		`u.id "user.id"`,
		`u.created_at "user.created_at"`,
		`u.updated_at "user.updated_at"`,
		`u.portal_code "user.employee.portal_code"`,
		`u.profile_id "user.profile_id"`,
		`u.email "user.email"`,
		`u.phone "user.phone"`,
		`u.role "user.role"`,
		`u.employee_first_name "user.employee.first_name"`,
		`u.employee_middle_name "user.employee.middle_name"`,
		`u.employee_last_name "user.employee.last_name"`,
		`u.employee_avatar "user.employee.avatar"`,
		`u.employee_position "user.employee.position"`,
		`u.employee_email "user.employee.email"`,
		`u.employee_phone "user.employee.phone"`,
	}

	return base.Builder().
		Select(selectColumns...).
		From(outmember.RoomOutmemberTableName).
		LeftJoin("users u on user_id = u.id").
		OrderBy("created_at desc").
		Where("address_id=?", roomID)
}

func roomSlotsSelectQuery(roomID uuid.UUID) sq.SelectBuilder {
	return base.Builder().
		Select(
			"id",
			"room_id",
			"timeline_start_at",
			"timeline_end_at",
			"plan_slot",
			"fact_slot",
			"created_at",
		).
		From(slotTableName).
		OrderBy("created_at desc").
		Where("room_id=?", roomID)
}

func roomEquipmentsSelectQuery(roomID uuid.UUID) sq.SelectBuilder {
	return base.Builder().
		Select(
			"id",
			"room_id",
			"title",
			"timeline_start_at",
			"timeline_end_at",
			"description",
			"plan_value",
			"fact_value",
			"created_at",
		).
		From(equipmentTableName).
		OrderBy("created_at desc").
		Where("room_id=?", roomID)
}

func roomReleaseSelectQuery(roomID uuid.UUID) sq.SelectBuilder {
	return base.Builder().
		Select(
			"id",
			"room_id",
			"title",
			"description",
			"date",
			"fact_slot",
			"created_at",
		).
		From(releaseTableName).
		OrderBy("created_at desc").
		Where("room_id=?", roomID)
}

func getUserColumns(tableTemplate, userTemplate string) []string {
	columns := []string{
		`tableTemplate.id "userTemplate.id"`,
		`tableTemplate.created_at "userTemplate.created_at"`,
		`tableTemplate.updated_at "userTemplate.updated_at"`,
		`tableTemplate.portal_code "userTemplate.portal_code"`,
		`tableTemplate.profile_id "userTemplate.profile_id"`,
		`tableTemplate.email "userTemplate.email"`,
		`tableTemplate.phone "userTemplate.phone"`,
		`tableTemplate.role "userTemplate.role"`,
		`tableTemplate.employee_first_name "userTemplate.employee_first_name"`,
		`tableTemplate.employee_middle_name "userTemplate.employee_middle_name"`,
		`tableTemplate.employee_last_name "userTemplate.employee_last_name"`,
		`tableTemplate.employee_avatar "userTemplate.employee_avatar"`,
		`tableTemplate.employee_position "userTemplate.employee_position"`,
		`tableTemplate.employee_email "userTemplate.employee_email"`,
		`tableTemplate.employee_phone "userTemplate.employee_phone"`,
	}

	for i, column := range columns {
		columns[i] = strings.ReplaceAll(strings.ReplaceAll(column, "userTemplate", userTemplate),
			"tableTemplate", tableTemplate)
	}

	return columns
}

func bookingWithUsersSelectQuery(bookingID []uuid.UUID) sq.SelectBuilder {
	return base.
		Builder().
		Select(
			"p.id",
			"p.number",
			"p.created_at",
			"p.updated_at",
			"p.title",
			"p.city",
			"p.description",
			"p.slot",
			"p.goal",
			"p.status",
			"p.links",
			"p.type",
			"p.state",
			`u.id "supervisor.id"`,
			`u.created_at "supervisor.created_at"`,
			`u.updated_at "supervisor.updated_at"`,
			`u.portal_code "supervisor.portal_code"`,
			`u.profile_id "supervisor.profile_id"`,
			`u.email "supervisor.email"`,
			`u.phone "supervisor.phone"`,
			`u.role "supervisor.role"`,
			`u.employee_first_name "supervisor.employee_first_name"`,
			`u.employee_middle_name "supervisor.employee_middle_name"`,
			`u.employee_last_name "supervisor.employee_last_name"`,
			`u.employee_avatar "supervisor.employee_avatar"`,
			`u.employee_email "supervisor.employee_email"`,
			`u.employee_phone "supervisor.employee_phone"`,
			`u.employee_position "supervisor.employee_position"`,
			`u1.id "author.id"`,
			`u1.created_at "author.created_at"`,
			`u1.updated_at "author.updated_at"`,
			`u1.portal_code "author.portal_code"`,
			`u1.profile_id "author.profile_id"`,
			`u1.email "author.email"`,
			`u1.phone "author.phone"`,
			`u1.role "author.role"`,
			`u1.employee_first_name "author.employee_first_name"`,
			`u1.employee_middle_name "author.employee_middle_name"`,
			`u1.employee_last_name "author.employee_last_name"`,
			`u1.employee_avatar "author.employee_avatar"`,
			`u1.employee_email "author.employee_email"`,
			`u1.employee_phone "author.employee_phone"`,
			`u1.employee_position "author.employee_position"`,
		).
		LeftJoin("booking_users book_supervisor on p.id = book_supervisor.booking_id and book_supervisor.role=1").
		LeftJoin("booking_users author on p.id = author.booking_id and author.role=3").
		LeftJoin("users u on book_supervisor.user_id = u.id").
		LeftJoin("users u1 on author.user_id = u1.id").
		From("bookings as p").
		Where(sq.Eq{"p.id": bookingID})
}

func slotsSelectQuery(slots []uuid.UUID) sq.SelectBuilder {
	return base.Builder().
		Select(
			"id",
			"room_id",
			"timeline_start_at",
			"timeline_end_at",
			"plan_slot",
			"fact_slot",
			"created_at",
		).
		From(slotTableName).
		Where(sq.Eq{"id": slots})
}

func equipmentsSelectQuery(equipments []uuid.UUID) sq.SelectBuilder {
	return base.Builder().
		Select(
			"id",
			"room_id",
			"title",
			"timeline_start_at",
			"timeline_end_at",
			"description",
			"plan_value",
			"fact_value",
			"created_at",
		).
		From(equipmentTableName).
		Where(sq.Eq{"id": equipments})
}

func releasesSelectQuery(releases []uuid.UUID) sq.SelectBuilder {
	return base.Builder().
		Select(
			"id",
			"room_id",
			"title",
			"description",
			"date",
			"fact_slot",
			"created_at",
		).
		From(releaseTableName).
		Where(sq.Eq{"id": releases})
}

func slotsByTimelineSelectQuery(roomID uuid.UUID, timeline model.Timeline) sq.SelectBuilder {
	return base.Builder().
		Select(
			"id",
			"room_id",
			"timeline_start_at",
			"timeline_end_at",
			"plan_slot",
			"fact_slot",
			"created_at",
		).
		From(slotTableName).
		Where("((timeline_start_at between ? and ?) or (timeline_end_at between ? and ?)) and room_id = ?",
			timeline.StartAt, timeline.EndAt, timeline.StartAt, timeline.EndAt, roomID)
}

func equipmentsTimelineSelectQuery(roomID uuid.UUID, timeline model.Timeline) sq.SelectBuilder {
	return base.Builder().
		Select(
			"id",
			"room_id",
			"title",
			"timeline_start_at",
			"timeline_end_at",
			"description",
			"plan_value",
			"fact_value",
			"created_at",
		).
		From(equipmentTableName).
		Where("((timeline_start_at between ? and ?) or (timeline_end_at between ? and ?)) and room_id = ?",
			timeline.StartAt, timeline.EndAt, timeline.StartAt, timeline.EndAt, roomID)
}

func releasesTimelineSelectQuery(roomID uuid.UUID, timeline model.Timeline) sq.SelectBuilder {
	return base.Builder().
		Select(
			"id",
			"room_id",
			"title",
			"description",
			"date",
			"fact_slot",
			"created_at",
		).
		From(releaseTableName).
		Where(
			"(date between ? and ?) and room_id = ?",
			timeline.StartAt, timeline.EndAt, roomID)
}

func slotSelectQuery(slotID uuid.UUID) sq.SelectBuilder {
	return base.
		Builder().
		Select(
			"id",
			"room_id",
			"timeline_start_at",
			"timeline_end_at",
			"plan_slot",
			"fact_slot",
			"created_at",
		).
		From(slotTableName).
		Where("id = ?", slotID)
}

func equipmentSelectQuery(equipmentID uuid.UUID) sq.SelectBuilder {
	return base.
		Builder().
		Select(
			"id",
			"room_id",
			"title",
			"timeline_start_at",
			"timeline_end_at",
			"description",
			"plan_value",
			"fact_value",
			"created_at",
		).
		From(equipmentTableName).
		Where("id = ?", equipmentID)
}

func releaseSelectQuery(releaseID uuid.UUID) sq.SelectBuilder {
	return base.Builder().
		Select(
			"id",
			"room_id",
			"title",
			"description",
			"date",
			"fact_slot",
			"created_at",
		).
		From(releaseTableName).
		Where("id = ?", releaseID)
}

func reportsSelectQuery(roomID uuid.UUID) sq.SelectBuilder {
	return base.
		Builder().
		Select(
			"id",
			"created_at",
			"updated_at",
			"room_id",
			"timeline_start_at",
			"timeline_end_at",
			"comment",
			"slots",
			"equipments",
			"releases",
		).
		From(reportTableName).
		Where("room_id = ?", roomID).
		OrderBy("created_at desc")
}

func roomParticipantsSelectQuery(roomID uuid.UUID) sq.SelectBuilder {
	return base.
		Builder().
		Select(
			"p.user_id",
			"p.role",
			`u.id "user.id"`,
			`u.created_at "user.created_at"`,
			`u.updated_at "user.updated_at"`,
			`u.portal_code "user.portal_code"`,
			`u.profile_id "user.profile_id"`,
			`u.email "user.email"`,
			`u.phone "user.phone"`,
			`u.role "user.role"`,
			`u.employee_first_name "user.employee_first_name"`,
			`u.employee_middle_name "user.employee_middle_name"`,
			`u.employee_last_name "user.employee_last_name"`,
			`u.employee_avatar "user.employee_avatar"`,
			`u.employee_email "user.employee_email"`,
			`u.employee_phone "user.employee_phone"`,
			`u.employee_position "user.employee_position"`,
		).
		From(
			fmt.Sprintf(
				`(select 
							(jsonb_array_elements(participants) ->> 'UserID')::uuid as user_id, 
							jsonb_array_elements(participants) ->> 'Role' as role 
						from rooms 
						WHERE rooms.id = '%s') p`,
				roomID),
		).
		LeftJoin("users u ON u.id = p.user_id")
}
