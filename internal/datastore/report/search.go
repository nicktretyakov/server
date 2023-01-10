package report

import (
	"context"

	"github.com/google/uuid"

	"be/internal/datastore/dbmodel"
	"be/internal/model"
)

func (s *Storage) ListByBookingID(ctx context.Context, bookingID uuid.UUID) ([]model.ReportBooking, error) {
	query := SelectQuery().Where("reports.booking_id=?", bookingID)

	reportsList := make(dbmodel.ReportBookingList, 0)
	if err := s.db.Select(ctx, query, &reportsList); err != nil {
		return nil, err
	}

	return reportsList.ToModel(), nil
}
