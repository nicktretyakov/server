package booking

import (
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"be/internal/datastore/filters"
	"be/internal/model"
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
			"(p.title ilike ? or p.description ilike ? or goal ilike ? or p.number::text ilike ?)", q, q, q, q,
		)
}

func NewSearchFilter(t string) filters.Filter {
	return SearchFilter{Query: t}
}

type StatusFilter struct {
	StatusList []model.Status
}

func (s *StatusFilter) Apply(builder sq.SelectBuilder) sq.SelectBuilder {
	return builder.Where(sq.Eq{"p.status": s.StatusList})
}

func NewStatusFilter(s []model.Status) filters.Filter {
	return &StatusFilter{
		StatusList: s,
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

type TimeLineFilter struct {
	To time.Time
}

func (s TimeLineFilter) Apply(builder sq.SelectBuilder) sq.SelectBuilder {
	return builder.Where("? between timeline_start_at and timeline_end_at", s.To)
}

func NewTimeLineFilter(to time.Time) filters.Filter {
	return TimeLineFilter{
		To: to,
	}
}

type SupervisorFilter struct{ UserID uuid.UUID }

func (s SupervisorFilter) Apply(builder sq.SelectBuilder) sq.SelectBuilder {
	return builder.Where(`(book_supervisor.user_id=?)`, s.UserID)
}

func NewSupervisorFilter(userID uuid.UUID) filters.Filter {
	return SupervisorFilter{
		UserID: userID,
	}
}

type SupervisorManyFilter struct{ UserIDs []uuid.UUID }

func (s SupervisorManyFilter) Apply(builder sq.SelectBuilder) sq.SelectBuilder {
	return builder.Where(sq.Eq{"book_supervisor.user_id": s.UserIDs})
}

func NewSupervisorManyFilter(userIDs []uuid.UUID) filters.Filter {
	return SupervisorManyFilter{
		UserIDs: userIDs,
	}
}

func NewAssigneeConfirmationWaitFilter(userID uuid.UUID) *AssigneeConfirmationWaitFilter {
	return &AssigneeConfirmationWaitFilter{userID: userID}
}

// AssigneeConfirmationWaitFilter filters
// userID - согласующий
// Проекты в статусе "model.OnAgreeAddressStatus"
// OR проекты в статусе "model.ConfirmedAddressStatus", есть завершающий отчёт в статусе model.OnAgreeFinalReportStatus.
type AssigneeConfirmationWaitFilter struct{ userID uuid.UUID }

func (a AssigneeConfirmationWaitFilter) Apply(builder sq.SelectBuilder) sq.SelectBuilder {
	filterString := `p.state=? and assignee.user_id=? and (p.status=? or p.status=?)`
	args := []interface{}{model.PublishedAddressState, a.userID, model.OnAgreeAddressStatus, model.FinalizeOnAgreeStatus}

	return builder.Where(filterString, args...)
}

// NewPOHeadConfirmationWaitFilter filters.
func NewPOHeadConfirmationWaitFilter() *POHeadConfirmationWaitFilter {
	return &POHeadConfirmationWaitFilter{}
}

type POHeadConfirmationWaitFilter struct{}

// Apply returns:
// Проекты в статусе "model.OnRegisterAddressStatus"
// OR проекты в статусе "model.ConfirmedAddressStatus", есть завершающий отчёт в статусе model.OnRegisterFinalReportStatus.
func (a POHeadConfirmationWaitFilter) Apply(builder sq.SelectBuilder) sq.SelectBuilder {
	filterString := `(p.status=? or p.status=?) and p.state=?`
	args := []interface{}{model.OnRegisterAddressStatus, model.FinalizeOnRegisterStatus, model.PublishedAddressState}

	return builder.Where(filterString, args...)
}

// DepartmentManyFilter filters.
type DepartmentManyFilter struct{ DepartmentIDs []model.DepartmentID }

func (s DepartmentManyFilter) Apply(builder sq.SelectBuilder) sq.SelectBuilder {
	conditions := make(sq.Or, 0)

	for _, departmentID := range s.DepartmentIDs {
		conditions = append(conditions, sq.Expr("(departments::jsonb @> ?) is true", []model.DepartmentID{departmentID}))
	}

	return builder.Where(conditions)
}

func NewDepartmentManyFilter(departmentIDs []model.DepartmentID) filters.Filter {
	return DepartmentManyFilter{
		DepartmentIDs: departmentIDs,
	}
}

// SlotFilter filters.
type SlotFilter struct {
	from *model.Notification
	to   *model.Notification
}

func (s SlotFilter) Apply(builder sq.SelectBuilder) sq.SelectBuilder {
	conditions := make(sq.And, 0)

	if s.from != nil {
		conditions = append(conditions, sq.Expr("p.slot >= ?", s.from.String()))
	}

	if s.to != nil {
		conditions = append(conditions, sq.Expr("p.slot <= ?", s.to.String()))
	}

	if len(conditions) == 0 {
		return builder
	}

	return builder.Where(conditions)
}

func NewSlotFilter(from *model.Notification, to *model.Notification) filters.Filter {
	return SlotFilter{
		from: from,
		to:   to,
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
