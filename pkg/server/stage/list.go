package stage

import (
	"context"

	"github.com/google/uuid"
	bookingpb "be/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"be/pkg/server/pbs"
)

func (s Service) GetListStages(ctx context.Context, req *bookingpb.GetListStagesRequest) (*bookingpb.GetListStagesResponse, error) {
	if req.GetBooking() == nil {
		return nil, status.New(codes.InvalidArgument, "missing booking aggregate").Err()
	}

	bookingID, err := uuid.Parse(req.Booking.GetUuid())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "wrong booking id").Err()
	}

	stageList, err := s.store.Stage().ListByBookingID(ctx, bookingID)
	if err != nil {
		return nil, err
	}

	return &bookingpb.GetListStagesResponse{
		Stages: pbs.StageList(stageList),
	}, nil
}
