package dbmodel

import (
	"time"

	"github.com/google/uuid"

	"be/internal/model"
)

type Stage struct {
	ID        uuid.UUID             `db:"id" yaml:"id"`
	BookingID uuid.UUID             `db:"booking_id" yaml:"booking_id"`
	Title     string                `db:"title" yaml:"title"`
	Status    model.AggregateStatus `db:"status" yaml:"status"`
	StartAt   *time.Time            `db:"start_at" yaml:"start_at"`
	EndAt     *time.Time            `db:"end_at" yaml:"end_at"`
	CreatedAt time.Time             `db:"created_at" yaml:"created_at"`
	Issues    []Issue               `db:"-" yaml:"-"`
}

func (r Stage) ToModel() model.Stage {
	s := model.Stage{
		ID:        r.ID,
		BookingID: r.BookingID,
		Title:     r.Title,
		CreatedAt: r.CreatedAt,
		Status:    r.Status,
	}

	if r.StartAt != nil && r.EndAt != nil {
		s.Timeline = &model.Timeline{
			StartAt: *r.StartAt,
			EndAt:   *r.EndAt,
		}
	}

	return s
}

func (r Stage) ToModelPtr() *model.Stage {
	repPtr := r.ToModel()
	return &repPtr
}

func StageFromModel(p model.Stage) Stage {
	s := Stage{
		ID:        p.ID,
		BookingID: p.BookingID,
		Title:     p.Title,
		CreatedAt: p.CreatedAt,
		Status:    p.Status,
	}

	if p.Timeline != nil {
		s.StartAt = &p.Timeline.StartAt
		s.EndAt = &p.Timeline.EndAt
	}

	return s
}

type StagesList []Stage

func (l StagesList) Stages() []model.Stage {
	modelsList := make([]model.Stage, 0, len(l))
	for _, stage := range l {
		modelsList = append(modelsList, stage.ToModel())
	}

	return modelsList
}

func (l StagesList) IDList() []uuid.UUID {
	idList := make([]uuid.UUID, 0, len(l))
	for _, stage := range l {
		idList = append(idList, stage.ID)
	}

	return idList
}
