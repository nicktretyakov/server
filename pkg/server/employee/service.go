package employee

import (
	"github.com/rs/zerolog"
	bookingpb "be/proto"

	"be/internal/profile"
)

type Service struct {
	bookingpb.UnsafeEmployeeServiceServer
	employeeRep profile.IProfile
	logger      zerolog.Logger
}

type Opts struct {
	EmployeeRep profile.IProfile
	Logger      zerolog.Logger
}

func New(opts Opts) bookingpb.EmployeeServiceServer {
	return &Service{logger: opts.Logger, employeeRep: opts.EmployeeRep}
}
