package dbmodel

import (
	"time"

	"github.com/google/uuid"

	"be/internal/model"
)

type IssueList []Issue

func (l IssueList) ToModels() []model.Issue {
	res := make([]model.Issue, 0, len(l))
	for _, issue := range l {
		res = append(res, issue.ToModel())
	}

	return res
}

func (l IssueList) IDList() []uuid.UUID {
	res := make([]uuid.UUID, 0, len(l))
	for _, issue := range l {
		res = append(res, issue.ID)
	}

	return res
}

type Issue struct {
	ID          uuid.UUID             `db:"id" yaml:"id"`
	Status      model.AggregateStatus `db:"status" yaml:"status"`
	StageID     uuid.UUID             `db:"stage_id" yaml:"stage_id"`
	Title       string                `db:"title" yaml:"title"`
	Description string                `db:"description" yaml:"description"`
	StartAt     *time.Time            `db:"start_at" yaml:"start_at"`
	EndAt       *time.Time            `db:"end_at" yaml:"end_at"`
	CreatedAt   time.Time             `db:"created_at" yaml:"created_at"`

	Participants UserList       `db:"participants"`
	Attachments  AttachmentList `db:"-"`
}

func (i Issue) ToModel() model.Issue {
	issue := model.Issue{
		ID:           i.ID,
		Status:       i.Status,
		Stage:        model.Stage{ID: i.StageID},
		CreatedAt:    i.CreatedAt,
		Title:        i.Title,
		Description:  i.Description,
		Participants: i.Participants.Users(),
		Attachments:  i.Attachments.Attachments(),
	}

	if i.StartAt != nil && i.EndAt != nil {
		issue.Timeline = &model.Timeline{
			StartAt: *i.StartAt,
			EndAt:   *i.EndAt,
		}
	}

	return issue
}

func (i Issue) ToModelPtr() *model.Issue {
	m := i.ToModel()
	return &m
}

func IssueFromModel(issue model.Issue) Issue {
	i := Issue{
		ID:           issue.ID,
		Status:       issue.Status,
		StageID:      issue.Stage.ID,
		Title:        issue.Title,
		Description:  issue.Description,
		CreatedAt:    issue.CreatedAt,
		Participants: UserListFromModel(issue.Participants),
		Attachments:  AttachmentListFromModel(issue.Attachments),
	}
	if issue.Timeline != nil {
		i.StartAt = &issue.Timeline.StartAt
		i.EndAt = &issue.Timeline.EndAt
	}

	return i
}
