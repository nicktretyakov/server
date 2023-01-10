package finalreports

import (
	"context"

	"github.com/google/uuid"

	"be/internal/datastore/dbmodel"
	"be/internal/model"
)

func (s *storage) FindByBookingID(ctx context.Context, bookingID uuid.UUID) (*model.FinalReport, error) {
	query := SelectQuery().Where("final_reports.booking_id=?", bookingID)

	var rep dbmodel.FinalReport
	if err := s.db.Get(ctx, query, &rep); err != nil {
		return nil, err
	}

	return rep.ToModelPtr(), nil
}

func (s *storage) FindByID(ctx context.Context, reportID uuid.UUID) (*model.FinalReport, error) {
	query := SelectQuery().Where("final_reports.id=?", reportID)

	var rep dbmodel.FinalReport
	if err := s.db.Get(ctx, query, &rep); err != nil {
		return nil, err
	}

	return rep.ToModelPtr(), nil
}
