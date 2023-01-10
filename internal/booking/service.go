package booking

import (
	"be/internal/datastore"
	"be/internal/report"
	"be/internal/user"
)

type Service struct {
	store         datastore.IDatastore
	userService   user.IUserService
	reportService report.IReportService
}

func New(store datastore.IDatastore, userService user.IUserService, reportService report.IReportService, opts ...Option) *Service {
	s := &Service{
		store:         store,
		userService:   userService,
		reportService: reportService,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

type Option func(s *Service)
