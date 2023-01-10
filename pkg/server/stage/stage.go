package stage

import (
	"context"
	"time"

	"github.com/google/uuid"
	bookingpb "be/proto"

	"be/internal/model"
	"be/pkg/server/pbs"
)

func (s Service) stageFromDB(ctx context.Context, stageID uuid.UUID) (*bookingpb.Stage, error) {
	p, err := s.store.Stage().FindByID(ctx, stageID)
	if err != nil {
		return nil, err
	}

	return pbs.Stage(*p), nil
}

func getTimelineModel(timeline *bookingpb.Timeline) (*model.Timeline, error) {
	start, err := time.Parse(pbs.TimeLayout, timeline.GetStart())
	if err != nil {
		return nil, err
	}

	end, err := time.Parse(pbs.TimeLayout, timeline.GetEnd())
	if err != nil {
		return nil, err
	}

	return &model.Timeline{
		StartAt: start,
		EndAt:   end,
	}, nil
}
