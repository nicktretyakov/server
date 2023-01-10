package pbs

import (
	bookingpb "be/proto"

	"be/internal/lib"
	"be/internal/model"
)

func outmember(agr model.Outmember) *bookingpb.Outmember {
	return &bookingpb.Outmember{
		Id:        agr.ID.String(),
		Object:    bookingpb.ObjectType(agr.Type),
		CreatedAt: ToUTCString(agr.CreatedAt),
		Result:    agr.Result.Bool(),
		Comment: &bookingpb.Comment{
			User:  PbUser(agr.User),
			Role:  bookingpb.Role(agr.Role),
			Extra: string(lib.MustJSON(agr.Extra)),
		},
	}
}

func OutmembersList(outmembers []model.Outmember) []*bookingpb.Outmember {
	res := make([]*bookingpb.Outmember, 0, len(outmembers))
	for _, agr := range outmembers {
		res = append(res, outmember(agr))
	}

	return res
}
