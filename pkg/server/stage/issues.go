package stage

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	bookingpb "be/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"be/pkg/auth"
	"be/pkg/server/pbs"
)

func (s Service) CreateIssue(ctx context.Context, req *bookingpb.CreateIssueRequest) (*bookingpb.IssueIDMessage, error) {
	stageID, err := uuid.Parse(req.GetStageId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid stage_id")
	}

	user := auth.FromContext(ctx)

	createdIssue, err := s.bookingService.AddBookingStageIssue(ctx, *user, stageID)
	if err != nil {
		return nil, err
	}

	return &bookingpb.IssueIDMessage{IssueId: createdIssue.ID.String()}, nil
}

func (s Service) UpdateIssue(ctx context.Context, req *bookingpb.UpdateIssueRequest) (*bookingpb.Issue, error) {
	issueToUpdate, err := validatedIssueRequest(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	issueToUpdate.ID, err = uuid.Parse(req.GetId())
	if err != nil {
		return nil, errors.New("invalid issue_id")
	}

	user := auth.FromContext(ctx)

	issue, err := s.bookingService.UpdateBookingStageIssue(ctx, *user, issueToUpdate)
	if err != nil {
		return nil, err
	}

	return pbs.Issue(*issue), nil
}

func (s Service) RemoveIssue(ctx context.Context, req *bookingpb.IssueIDMessage) (*emptypb.Empty, error) {
	issueID, err := uuid.Parse(req.GetIssueId())
	if err != nil {
		return nil, errors.New("invalid issue_id")
	}

	user := auth.FromContext(ctx)

	issue, err := s.store.Stage().FindStageIssueByID(ctx, issueID)
	if err != nil {
		return nil, err
	}

	if err = s.bookingService.RemoveBookingStageIssue(ctx, *user, *issue); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
