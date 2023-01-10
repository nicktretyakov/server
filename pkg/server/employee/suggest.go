package employee

import (
	"context"

	bookingpb "be/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"be/pkg/server/pbs"
)

const minQueryLength = 3

func (s Service) SuggestEmployee(ctx context.Context, req *bookingpb.SuggestEmployeeRequest) (*bookingpb.SuggestEmployeeResponse, error) {
	query := req.GetQuery()
	if len(query) < minQueryLength {
		return nil, status.Errorf(codes.InvalidArgument, "query min length=%d", minQueryLength)
	}

	limit := req.GetLimit()
	offset := req.GetOffset()

	profileEmployees, err := s.employeeRep.FindEmployees(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}

	return &bookingpb.SuggestEmployeeResponse{
		Employees: pbs.PbEmployeeList(profileEmployees.Cast()),
	}, nil
}
