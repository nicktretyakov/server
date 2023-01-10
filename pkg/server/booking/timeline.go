package booking

import (
	"time"

	bookingpb "be/proto"

	"be/internal/model"
	"be/pkg/server/pbs"
)

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
