package stage

import (
	"context"

	"github.com/google/uuid"

	"be/internal/datastore/dbmodel"
	"be/internal/model"
)

func (s *Storage) FindByID(ctx context.Context, stageID uuid.UUID) (*model.Stage, error) {
	selectQuery := selectQuery().Where("stages.id=?", stageID)

	var rep dbmodel.Stage
	if err := s.db.Get(ctx, selectQuery, &rep); err != nil {
		return nil, err
	}

	issues, err := s.findStagesIssues(ctx, []uuid.UUID{stageID})
	if err != nil {
		return nil, err
	}

	stage := rep.ToModel()
	stage.Issues = issues

	return &stage, nil
}

func (s Storage) ListByBookingID(ctx context.Context, bookingID uuid.UUID) ([]model.Stage, error) {
	query := selectQuery().Where("stages.booking_id=? and status=?", bookingID, model.ActiveAggregateStatus)

	stagesList := make(dbmodel.StagesList, 0)
	if err := s.db.Select(ctx, query, &stagesList); err != nil {
		return nil, err
	}

	issues, err := s.findStagesIssues(ctx, stagesList.IDList())
	if err != nil {
		return nil, err
	}

	stageToIssueMap := make(map[uuid.UUID][]model.Issue)
	for _, issue := range issues {
		if _, ok := stageToIssueMap[issue.Stage.ID]; !ok {
			stageToIssueMap[issue.Stage.ID] = make([]model.Issue, 0, 1)
		}

		stageToIssueMap[issue.Stage.ID] = append(stageToIssueMap[issue.Stage.ID], issue)
	}

	stages := stagesList.Stages()

	for i, stage := range stages {
		if stageIssues, ok := stageToIssueMap[stage.ID]; ok {
			stages[i].Issues = stageIssues
		}
	}

	return stages, nil
}
