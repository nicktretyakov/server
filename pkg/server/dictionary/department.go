package dictionary

import (
	"context"

	bookingpb "be/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"be/pkg/server/pbs"
)

func (s Service) DepartmentList(ctx context.Context, _ *emptypb.Empty) (*bookingpb.DepartmentListResponse, error) {
	deps, err := s.store.Dictionary().DepartmentList(ctx)
	if err != nil {
		return nil, err
	}

	return &bookingpb.DepartmentListResponse{
		Departments: pbs.ToPbDepartments(deps),
	}, nil
}
