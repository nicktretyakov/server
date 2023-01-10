package stage

import (
	"github.com/rs/zerolog"
	bookingpb "be/proto"

	"be/internal/datastore"
	"be/internal/booking"
)

type Service struct {
	bookingpb.UnsafeStageServiceServer
	store          datastore.IDatastore
	bookingService booking.IBookingService
	logger         zerolog.Logger
}

type Opts struct {
	Store          datastore.IDatastore
	BookingService booking.IBookingService
	Logger         zerolog.Logger
}

func NewService(opts Opts) bookingpb.StageServiceServer {
	return &Service{store: opts.Store, logger: opts.Logger, bookingService: opts.BookingService}
}
