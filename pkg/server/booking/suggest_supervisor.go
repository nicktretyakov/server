package booking

import (
	"context"

	bookingpb "be/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"be/pkg/server/pbs"
)

func (s Service) SuggestSupervisor(ctx context.Context, request *emptypb.Empty) (*bookingpb.SuggestSupervisorResponse, error) {
	usersIDs, err := s.store.Booking().GetListSupervisors(ctx)
	if err != nil {
		return nil, err
	}

	users, err := s.store.User().FindUsersByIDs(ctx, usersIDs)
	if err != nil {
		return nil, err
	}

	return &bookingpb.SuggestSupervisorResponse{Users: pbs.PbUserList(users)}, nil
}
