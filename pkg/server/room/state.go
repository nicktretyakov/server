package room

import (
	"context"
	bookingpb "be/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"be/internal/acl"
	"be/internal/model"
	"be/pkg/auth"
)

func (s Service) SetRoomArchived(ctx context.Context, req *bookingpb.SetRoomArchivedRequest) (*emptypb.Empty, error) {
	if err := s.setState(ctx, req.GetRoomID(), model.ArchivedAddressState); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s Service) SetRoomPublished(ctx context.Context, req *bookingpb.SetRoomPublishedRequest) (*emptypb.Empty, error) {
	if err := s.setState(ctx, req.GetRoomID(), model.PublishedAddressState); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s Service) setState(ctx context.Context, rooID string, toState model.State) error {
	user := auth.FromContext(ctx)

	roomID, err := uuid.Parse(rooID)
	if err != nil {
		return status.New(codes.InvalidArgument, "invalid room_id").Err()
	}

	room, err := s.store.Room().FindByID(ctx, roomID)
	if err != nil {
		return err
	}

	if !acl.CanChangeState(*user, room) {
		return status.New(codes.PermissionDenied, "permission denied").Err()
	}

	if err = s.store.Room().UpdateState(ctx, roomID, toState); err != nil {
		return err
	}

	return nil
}
