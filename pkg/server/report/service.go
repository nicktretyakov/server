package report

import (
	"github.com/rs/zerolog"
	bookingpb "be/proto"

	"be/internal/datastore"
	"be/internal/notecreator"
	"be/internal/booking"
)

type Service struct {
	bookingpb.UnsafeReportServiceServer
	store          datastore.IDatastore
	logger         zerolog.Logger
	bookingService booking.IBookingService
	notificator    *notecreator.NoteCreator
}

type Opts struct {
	Store          datastore.IDatastore
	Logger         zerolog.Logger
	BookingService booking.IBookingService
	Notificator    *notecreator.NoteCreator
}

func NewService(opts Opts) bookingpb.ReportServiceServer {
	return &Service{store: opts.Store, logger: opts.Logger, bookingService: opts.BookingService, notificator: opts.Notificator}
}
