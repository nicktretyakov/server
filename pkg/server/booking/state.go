package booking

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

func (s Service) SetBookingArchived(ctx context.Context, req *bookingpb.SetBookingArchivedRequest) (*emptypb.Empty, error) {
	if err := s.setState(ctx, req.GetBookingID(), model.ArchivedAddressState); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s Service) SetBookingPublished(ctx context.Context, req *bookingpb.SetBookingPublishedRequest) (*emptypb.Empty, error) {
	if err := s.setState(ctx, req.GetBookingID(), model.PublishedAddressState); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s Service) setState(ctx context.Context, bookID string, toState model.State) error {
	user := auth.FromContext(ctx)

	bookingID, err := uuid.Parse(bookID)
	if err != nil {
		return status.New(codes.InvalidArgument, "invalid booking_id").Err()
	}

	book, err := s.store.Booking().FindByID(ctx, bookingID)
	if err != nil {
		return err
	}

	if !acl.CanChangeState(*user, book) {
		return status.New(codes.PermissionDenied, "permission denied").Err()
	}

	if err = s.store.Booking().UpdateState(ctx, bookingID, toState); err != nil {
		return err
	}

	return nil
}
