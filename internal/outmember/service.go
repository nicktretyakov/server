package outmember

import (
	"be/internal/datastore"
	"be/internal/report"
)

type Service struct {
	store         datastore.IDatastore
	reportService report.IReportService
}

func New(store datastore.IDatastore, reportService report.IReportService) *Service {
	return &Service{store: store, reportService: reportService}
}
