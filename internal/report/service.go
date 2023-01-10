package report

import (
	"context"

	"github.com/google/uuid"

	"be/internal/datastore"
	"be/internal/model"
)

type IReportService interface {
	CheckReports(ctx context.Context, bookingID *uuid.UUID, periods []model.Period) error
}

type service struct {
	store datastore.IReportStore
}

func New(store datastore.IReportStore) *service {
	return &service{store: store}
}
