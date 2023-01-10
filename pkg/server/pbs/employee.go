package pbs

import (
	bookingpb "be/proto"

	"be/internal/model"
)

func PbEmployee(e model.Employee) *bookingpb.Employee {
	return &bookingpb.Employee{
		Name:       e.GetFirstName(),
		Middlename: e.GetMiddleName(),
		Lastname:   e.GetLastName(),
		Avatar:     e.GetAvatarLink(),
		Email:      e.GetEmail(),
		Phone:      e.GetPhone(),
		Portalcode: uint32(e.PortalCode),
		Position:   e.GetPosition(),
	}
}

func PbEmployeeList(list []model.Employee) []*bookingpb.Employee {
	res := make([]*bookingpb.Employee, 0, len(list))

	for _, employee := range list {
		res = append(res, PbEmployee(employee))
	}

	return res
}
