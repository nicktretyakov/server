package stage

import (
	"context"

	bookingpb "be/proto"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"be/internal/acl"
	"be/internal/model"
	"be/pkg/auth"
)

func (s Service) UpdateStage(ctx context.Context, req *bookingpb.UpdateStageRequest) (*bookingpb.UpdateStageResponse, error) {
	validatedStage, err := ValidateUpdateStageRequest(req)
	if err != nil {
		return nil, err
	}

	stageToUpdate, err := s.store.Stage().FindByID(ctx, validatedStage.ID)
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "cannot update stage").Err()
	}

	booking, err := s.store.Booking().FindByID(ctx, stageToUpdate.BookingID)
	if err != nil {
		return nil, err
	}

	user := auth.FromContext(ctx)
	if !acl.CanUpdateStage(*user, *booking) {
		return nil, status.New(codes.PermissionDenied, "not enough privileges").Err()
	}

	validatedStage.update(stageToUpdate)
	stageToUpdate.Status = model.ActiveAggregateStatus

	_, err = s.store.Stage().Update(ctx, *stageToUpdate)
	if err != nil {
		return nil, err
	}

	pbStage, err := s.stageFromDB(ctx, validatedStage.ID)
	if err != nil {
		return nil, err
	}

	return &bookingpb.UpdateStageResponse{Stage: pbStage}, nil
}

func ValidateUpdateStageRequest(req *bookingpb.UpdateStageRequest) (*updateRequest, error) {
	if req.GetStage() == nil {
		return nil, status.New(codes.InvalidArgument, "missing stage aggregate").Err()
	}

	stageID, err := uuid.Parse(req.Stage.GetUuid())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "wrong stage id").Err()
	}

	if req.Stage.GetTimeline() == nil {
		return nil, status.New(codes.InvalidArgument, "missing stage timeline aggregate").Err()
	}

	title := req.Stage.GetTitle()
	if len(title) < MinTitleLength {
		return nil, status.Errorf(codes.InvalidArgument, "invalid stage name (min. length %d characters)", MinTitleLength)
	}

	timeline, err := getTimelineModel(req.Stage.GetTimeline())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "invalid date format").Err()
	}

	if !timeline.IsValid() {
		return nil, status.New(codes.InvalidArgument, "invalid timeline").Err()
	}

	return &updateRequest{
		ID:       stageID,
		Title:    title,
		Timeline: *timeline,
	}, nil
}

type updateRequest struct {
	ID       uuid.UUID
	Title    string
	Timeline model.Timeline
}

func (r updateRequest) update(book *model.Stage) {
	book.Title = r.Title
	book.Timeline = &r.Timeline
}
