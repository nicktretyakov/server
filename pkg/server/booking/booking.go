package booking

import (
	"context"

	"github.com/google/uuid"
	bookingpb "be/proto"

	"be/pkg/server/pbs"
)

func (s Service) bookingFromDB(ctx context.Context, bookingID uuid.UUID) (*bookingpb.Booking, error) {
	p, err := s.store.Booking().FindByID(ctx, bookingID)
	if err != nil {
		return nil, err
	}

	return pbs.PbBooking(*p), nil
}
