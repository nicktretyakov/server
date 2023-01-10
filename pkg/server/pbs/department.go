package pbs

import (
	bookingpb "be/proto"

	"be/internal/model"
)

func ToPbDepartments(deps []model.Department) []*bookingpb.Department {
	pbDeps := make([]*bookingpb.Department, 0, len(deps))
	for _, dep := range deps {
		pbDeps = append(pbDeps, &bookingpb.Department{
			Id:   dep.ID.String(),
			Name: dep.Title,
		})
	}

	return pbDeps
}
