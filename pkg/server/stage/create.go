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

const (
	MinTitleLength = 3
)

func (s Service) CreateStage(ctx context.Context, req *bookingpb.CreateStageRequest) (*bookingpb.CreateStageResponse, error) {
	validatedStage, err := ValidateCreateStageRequest(req)
	if err != nil {
		return nil, err
	}

	booking, err := s.store.Booking().FindByID(ctx, validatedStage.BookingID)
	if err != nil {
		return nil, err
	}

	user := auth.FromContext(ctx)
	if !acl.CanAddStage(*user, *booking) {
		return nil, status.New(codes.PermissionDenied, "not enough privileges").Err()
	}

	validatedStage.Status = model.InitialAggregateStatus

	createdStage, err := s.store.Stage().Create(ctx, *validatedStage)
	if err != nil {
		return nil, err
	}

	return &bookingpb.CreateStageResponse{
		StageId: createdStage.ID.String(),
	}, nil
}

func ValidateCreateStageRequest(req *bookingpb.CreateStageRequest) (*model.Stage, error) {
	bookingID, err := uuid.Parse(req.GetBookingId())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "wrong booking id").Err()
	}

	return &model.Stage{BookingID: bookingID}, nil
}
