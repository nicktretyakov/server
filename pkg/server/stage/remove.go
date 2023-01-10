package stage

import (
	"context"

	"github.com/google/uuid"
	bookingpb "be/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"be/internal/acl"
	"be/internal/model"
	"be/pkg/auth"
)

func (s Service) RemoveStage(ctx context.Context, req *bookingpb.RemoveStageRequest) (*bookingpb.RemoveStageResponse, error) {
	validatedStage, err := ValidateRemoveStageRequest(req)
	if err != nil {
		return nil, err
	}

	stageToRemove, err := s.store.Stage().FindByID(ctx, validatedStage.ID)
	if err != nil {
		return nil, err
	}

	booking, err := s.store.Booking().FindByID(ctx, stageToRemove.BookingID)
	if err != nil {
		return nil, err
	}

	user := auth.FromContext(ctx)
	if !acl.CanRemoveStage(*user, *booking) {
		return nil, status.New(codes.PermissionDenied, "not enough privileges").Err()
	}

	err = s.store.Stage().Remove(ctx, *stageToRemove)
	if err != nil {
		return nil, err
	}

	return &bookingpb.RemoveStageResponse{StageId: stageToRemove.ID.String()}, nil
}

func ValidateRemoveStageRequest(req *bookingpb.RemoveStageRequest) (*model.Stage, error) {
	if req.GetStage() == nil {
		return nil, status.New(codes.InvalidArgument, "missing stage aggregate").Err()
	}

	stageID, err := uuid.Parse(req.Stage.GetUuid())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "wrong stage id").Err()
	}

	return &model.Stage{ID: stageID}, nil
}
