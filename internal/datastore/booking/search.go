package booking

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"be/internal/datastore/base"
	"be/internal/datastore/dbmodel"
	"be/internal/datastore/filters"
	"be/internal/datastore/outmember"
	"be/internal/datastore/report"
	"be/internal/datastore/sorting"
	"be/internal/model"
)

// FindByID return booking with supervisor, departments and reports.
func (s *Storage) FindByID(ctx context.Context, bookingID uuid.UUID) (*model.Booking, error) {
	query := bookingSelectQuery().Where("p.id=?", bookingID)

	var book dbmodel.Booking
	if err := s.db.Get(ctx, query, &book); err != nil {
		return nil, err
	}

	bookReports := make(dbmodel.ReportBookingList, 0)
	if err := s.db.Select(ctx, bookingReportsSelectQuery(bookingID), &bookReports); err != nil {
		return nil, err
	}

	book.Reports = bookReports

	var bookOutmembers []model.Outmember
	if err := s.db.Select(ctx, bookingOutmembersSelectQuery(bookingID), &bookOutmembers); err != nil && !errors.Is(err, base.ErrNotFound) {
		return nil, err
	}

	book.Outmembers = bookOutmembers

	var bookingUsers []*dbmodel.BookingUser
	if err := s.db.Select(ctx, bookingUsersSelectQuery(bookingID), &bookingUsers); err != nil {
		return nil, err
	}

	book.SetUsers(bookingUsers)

	bookAttachments, err := s.getAttachmentsModelDB(ctx, bookingID)
	if err != nil {
		return nil, err
	}

	book.Attachments = bookAttachments

	return book.ToModelPtr(), nil
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

// ActiveList return booking with supervisor, final_report.
// Adds filter status!=1.
func (s *Storage) ActiveList(ctx context.Context, limit, offset uint64, sorting sorting.Sorting, queryFilters ...filters.Filter) ([]model.Booking, uint64, error) { //nolint:lll
	queryFilters = append(queryFilters, newStatusNegFilter([]model.Status{model.InitialAddressStatus}))
	query, countQuery := withFilters(limit, offset, queryFilters...)
	query = sorting.Apply(query)

	var (
		bookings = make(dbmodel.BookingList, 0)
		count    uint64
	)

	if err := s.db.Select(ctx, query, &bookings); err != nil {
		return nil, 0, err
	}

	if err := s.db.Get(ctx, countQuery, &count); err != nil {
		return nil, 0, err
	}

	for _, book := range bookings {
		var bookingUsers []*dbmodel.BookingUser

		if err := s.db.Select(ctx, bookingUsersParticipantsSelectQuery(book.ID), &bookingUsers); err != nil {
			return nil, 0, err
		}

		book.SetUsers(bookingUsers)
	}

	return bookings.Bookings(), count, nil
}

func (s *Storage) List(ctx context.Context) ([]model.Booking, uint64, error) {
	query, countQuery := withFilters(0, 0)

	var (
		bookings = make(dbmodel.BookingList, 0)
		count    uint64
	)

	if err := s.db.Select(ctx, query, &bookings); err != nil {
		return nil, 0, err
	}

	if err := s.db.Get(ctx, countQuery, &count); err != nil {
		return nil, 0, err
	}

	return bookings.Bookings(), count, nil
}

func (s *Storage) ListBookingsWithNotSendReports(ctx context.Context) ([]model.Booking, error) {
	var (
		bookings = make(dbmodel.BookingList, 0)
		query    = notSendReportsQuery(bookingWithUsersSelectQuery())
	)

	if err := s.db.Select(ctx, query, &bookings); err != nil {
		return nil, err
	}

	for _, booking := range bookings {
		bookingReports := make(dbmodel.ReportBookingList, 0)
		if err := s.db.Select(ctx, bookingReportsSelectQuery(booking.ID), &bookingReports); err != nil {
			return nil, err
		}

		booking.Reports = bookingReports
	}

	return bookings.Bookings(), nil
}

func withFilters(limit, offset uint64, queryFilters ...filters.Filter) (sq.SelectBuilder, sq.SelectBuilder) {
	query := bookingWithUsersSelectQuery()
	countQuery := bookingCountQuery()

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

func bookingCountQuery() sq.SelectBuilder {
	return base.Builder().Select("count(*)").
		LeftJoin("booking_users book_supervisor on p.id = book_supervisor.booking_id and book_supervisor.role=1").
		LeftJoin("booking_users assignee on p.id = assignee.booking_id and assignee.role=2").
		LeftJoin("booking_final_reports final_reports on final_reports.booking_id=p.id").
		LeftJoin(`(select pd.booking_id, json_agg(json_build_object('id', dd.id, 'title', dd.title)) departments
               from booking_departments pd join dictionary_departments dd on pd.department_id = dd.id
               group by booking_id) x on x.booking_id = p.id`).
		From(bookingTableNameAlias)
}

func bookingSelectQuery() sq.SelectBuilder {
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
			"p.timeline_start_at",
			"p.timeline_end_at",
			"p.status",
			"p.type",
			"p.links",
			"p.state",
			`x.departments "departments"`,
			`final_reports.id "final_report.id"`,
			`final_reports.created_at "final_report.created_at"`,
			`final_reports.updated_at "final_report.updated_at"`,
			`final_reports.booking_id "final_report.booking_id"`,
			`final_reports.slot "final_report.slot"`,
			`final_reports.end_at "final_report.end_at"`,
			`final_reports.comment "final_report.comment"`,
			`final_reports.status "final_report.status"`,
			`final_reports.attachments_uuid "final_report.attachments_uuid"`,
		).
		LeftJoin(`(select pd.booking_id, json_agg(json_build_object('id', dd.id, 'title', dd.title)) departments
               from booking_departments pd join dictionary_departments dd on pd.department_id = dd.id
               group by booking_id) x on x.booking_id = p.id`).
		LeftJoin("booking_final_reports final_reports on final_reports.booking_id=p.id").
		OrderBy("created_at desc").
		From(bookingTableNameAlias)
}

func bookingWithUsersSelectQuery() sq.SelectBuilder {
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
			"p.timeline_start_at",
			"p.timeline_end_at",
			"p.status",
			"p.links",
			"p.type",
			"p.state",
			`x.departments "departments"`,
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
			`final_reports.id "final_report.id"`,
			`final_reports.created_at "final_report.created_at"`,
			`final_reports.updated_at "final_report.updated_at"`,
			`final_reports.booking_id "final_report.booking_id"`,
			`final_reports.slot "final_report.slot"`,
			`final_reports.end_at "final_report.end_at"`,
			`final_reports.comment "final_report.comment"`,
			`final_reports.status "final_report.status"`,
			`final_reports.attachments_uuid "final_report.attachments_uuid"`,
		).
		LeftJoin("booking_users book_supervisor on p.id = book_supervisor.booking_id and book_supervisor.role=1").
		LeftJoin("booking_users assignee on p.id = assignee.booking_id and assignee.role=2").
		LeftJoin("users u on book_supervisor.user_id = u.id").
		LeftJoin(`(select pd.booking_id, json_agg(json_build_object('id', dd.id, 'title', dd.title)) departments
               from booking_departments pd join dictionary_departments dd on pd.department_id = dd.id
               group by booking_id) x on x.booking_id = p.id`).
		LeftJoin("booking_final_reports final_reports on final_reports.booking_id=p.id").
		From(bookingTableNameAlias)
}

func bookingReportsSelectQuery(bookingID uuid.UUID) sq.SelectBuilder {
	return report.SelectQuery().
		Where("reports.booking_id=? and date_trunc('month', reports.period) <= current_date", bookingID).
		OrderBy("reports.period desc")
}

func notSendReportsQuery(builder sq.SelectBuilder) sq.SelectBuilder {
	notSendReportsSQL := fmt.Sprintf("EXISTS(select * from %s WHERE period <= current_date AND status = %d AND booking_id = p.id)",
		bookingReportsTableName, model.NotSendReportStatus)

	return builder.Where("(p.status!=? and p.status!=?) and p.state=?",
		model.DoneAddressStatus,
		model.InitialAddressStatus,
		model.PublishedAddressState,
	).
		Where(notSendReportsSQL)
}

func bookingOutmembersSelectQuery(bookingID uuid.UUID) sq.SelectBuilder {
	return base.Builder().
		Select(
			`booking_outmember.id "id"`,
			`booking_outmember.created_at "created_at"`,
			`booking_outmember.type "type"`,
			`booking_outmember.user_id "user_id"`,
			`booking_outmember.address_id "address_id"`,
			`booking_outmember.result "result"`,
			`booking_outmember.extra "extra"`,
			`booking_outmember.role "role"`,
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
			`u.employee_phone "user.employee.phone"`).
		From(outmember.BookingOutmemberTableName).
		LeftJoin("users u on user_id = u.id").
		OrderBy("created_at desc").
		Where("address_id=?", bookingID)
}

func bookingUsersSelectQuery(bookingID uuid.UUID) sq.SelectBuilder {
	return base.Builder().
		Select(
			`booking_users.booking_id "booking_id"`,
			`booking_users.role "role"`,
			`booking_users.role_description "role_description"`,
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
			`u.employee_position "user.employee_position"`,
			`u.employee_email "user.employee_email"`,
			`u.employee_phone "user.employee_phone"`).
		From(bookingUserRoleTableName).
		LeftJoin("users u on user_id = u.id").
		Where("booking_id=?", bookingID)
}

func bookingUsersParticipantsSelectQuery(bookingID uuid.UUID) sq.SelectBuilder {
	return base.Builder().
		Select(
			`booking_users.booking_id "booking_id"`,
			`booking_users.role "role"`,
			`booking_users.role_description "role_description"`,
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
			`u.employee_position "user.employee_position"`,
			`u.employee_email "user.employee_email"`,
			`u.employee_phone "user.employee_phone"`).
		From(bookingUserRoleTableName).
		LeftJoin("users u on user_id = u.id").
		Where("booking_id=? and booking_users.role=?", bookingID, model.ParticipantBookingUserRole)
}
