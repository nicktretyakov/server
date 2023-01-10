package room

import (
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"be/internal/datastore/filters"
	"be/internal/model"
	"be/pkg/server/pbs"
)

type limitFilter struct{ limit uint64 }

func (l limitFilter) Apply(builder sq.SelectBuilder) sq.SelectBuilder {
	return builder.Limit(l.limit)
}

func newLimitFilter(limit uint64) filters.Filter {
	return limitFilter{limit: limit}
}

type OffsetFilter struct{ offset uint64 }

func (f OffsetFilter) Apply(builder sq.SelectBuilder) sq.SelectBuilder {
	return builder.Offset(f.offset)
}

func newOffsetFilter(offset uint64) filters.Filter {
	return OffsetFilter{offset: offset}
}

type SearchFilter struct {
	Query string
}

func (s SearchFilter) Apply(builder sq.SelectBuilder) sq.SelectBuilder {
	q := fmt.Sprint("%", s.Query, "%")

	return builder.
		Where(
			"(p.title ilike ? or p.description ilike ? or p.number::text ilike ?)", q, q, q,
		)
}

func NewSearchFilter(t string) filters.Filter {
	return SearchFilter{Query: t}
}

type AuthorFilter struct {
	UserID uuid.UUID
}

func (f AuthorFilter) Apply(builder sq.SelectBuilder) sq.SelectBuilder {
	return builder.Where("author_id = ?", f.UserID)
}

func NewAuthorFilter(authorID uuid.UUID) *AuthorFilter {
	return &AuthorFilter{
		UserID: authorID,
	}
}

type OwnerFilter struct {
	UserIDs []uuid.UUID
}

func (o OwnerFilter) Apply(builder sq.SelectBuilder) sq.SelectBuilder {
	return builder.Where(sq.Eq{"owner_id": o.UserIDs})
}

func NewOwnerFilter(ownerIDs []uuid.UUID) (*OwnerFilter, error) {
	return &OwnerFilter{
		UserIDs: ownerIDs,
	}, nil
}

type EmployeeFilter struct {
	UserIDs []uuid.UUID
}

func (e EmployeeFilter) Apply(builder sq.SelectBuilder) sq.SelectBuilder {
	return builder.Where(sq.Eq{"employee_id": e.UserIDs})
}

func NewEmployeeFilter(employeeIDs []uuid.UUID) (*EmployeeFilter, error) {
	return &EmployeeFilter{
		UserIDs: employeeIDs,
	}, nil
}

type POHeadConfirmationWaitFilter struct{}

// Apply returns:
// Продукты в статусе "model.OnRegisterAddressStatus".
func (a POHeadConfirmationWaitFilter) Apply(builder sq.SelectBuilder) sq.SelectBuilder {
	return builder.Where("p.status=? and p.state=?", model.OnRegisterAddressStatus, model.PublishedAddressState)
}

// NewPOHeadConfirmationWaitFilter filters.
func NewPOHeadConfirmationWaitFilter() *POHeadConfirmationWaitFilter {
	return &POHeadConfirmationWaitFilter{}
}

type AssigneeConfirmationWaitFilter struct{ userID uuid.UUID }

// Apply AssigneeConfirmationWaitFilter filters
// userID - согласующий
// Продукты в статусе "model.OnAgreeAddressStatus".
func (a AssigneeConfirmationWaitFilter) Apply(builder sq.SelectBuilder) sq.SelectBuilder {
	return builder.Where("employee_id=? and p.status=? and p.state=?",
		a.userID,
		model.OnAgreeAddressStatus,
		model.PublishedAddressState,
	)
}

func NewAssigneeConfirmationWaitFilter(assigneeID uuid.UUID) *AssigneeConfirmationWaitFilter {
	return &AssigneeConfirmationWaitFilter{
		userID: assigneeID,
	}
}

type StatusNegFilter struct {
	StatusList []model.Status
}

func (s *StatusNegFilter) Apply(builder sq.SelectBuilder) sq.SelectBuilder {
	return builder.Where(sq.NotEq{"p.status": s.StatusList})
}

func newStatusNegFilter(s []model.Status) filters.Filter {
	return &StatusNegFilter{
		StatusList: s,
	}
}

type CreatedAtFilter struct {
	From time.Time
	To   time.Time
}

func (c CreatedAtFilter) Apply(builder sq.SelectBuilder) sq.SelectBuilder {
	return builder.Where("p.created_at between ? and ?", c.From, c.To)
}

func NewCreatedAtFilter(createdAtStart, createdAtEnd string) (*CreatedAtFilter, error) {
	startAt, endAt, err := parseTimeline(createdAtStart, createdAtEnd)
	if err != nil {
		return nil, err
	}

	return &CreatedAtFilter{
		From: *startAt,
		To:   *endAt,
	}, nil
}

type CreationDateFilter struct {
	From time.Time
	To   time.Time
}

func (c CreationDateFilter) Apply(builder sq.SelectBuilder) sq.SelectBuilder {
	return builder.Where("p.creation_date between ? and ?", c.From, c.To)
}

func NewCreationDateFilter(creationDateStart, creationDateEnd string) (*CreationDateFilter, error) {
	startAt, endAt, err := parseTimeline(creationDateStart, creationDateEnd)
	if err != nil {
		return nil, err
	}

	return &CreationDateFilter{
		From: *startAt,
		To:   *endAt,
	}, nil
}

type StatusFilter struct {
	Status []model.Status
}

func (s StatusFilter) Apply(builder sq.SelectBuilder) sq.SelectBuilder {
	return builder.Where(sq.Eq{"p.status": s.Status})
}

func NewStatusFilter(status []model.Status) *StatusFilter {
	return &StatusFilter{
		Status: status,
	}
}

type StateFilter struct {
	State []model.State
}

func (s StateFilter) Apply(builder sq.SelectBuilder) sq.SelectBuilder {
	return builder.Where(sq.Eq{"p.state": s.State})
}

func NewStateFilter(state []model.State) *StateFilter {
	return &StateFilter{
		State: state,
	}
}

func parseTimeline(startAtStr, endAtStr string) (*time.Time, *time.Time, error) {
	startAt, err := time.ParseInLocation(pbs.TimeLayout, startAtStr, time.UTC)
	if err != nil {
		return nil, nil, err
	}

	endAt, err := time.ParseInLocation(pbs.TimeLayout, endAtStr, time.UTC)
	if err != nil {
		return nil, nil, err
	}

	return &startAt, &endAt, nil
}
