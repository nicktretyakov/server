package room

import (
	"context"

	"github.com/google/uuid"
	bookingpb "be/proto"

	"be/pkg/server/pbs"
)

func (s Service) roomFromDB(ctx context.Context, roomID uuid.UUID) (*bookingpb.Room, error) {
	p, err := s.store.Room().FindByID(ctx, roomID)
	if err != nil {
		return nil, err
	}

	return pbs.PbRoom(*p), nil
}
