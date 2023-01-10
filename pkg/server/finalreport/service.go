package finalreport

import (
	"context"

	"github.com/rs/zerolog"
	bookingpb "be/proto"

	"be/internal/datastore"
	"be/internal/model"
	"be/internal/notecreator"
	"be/internal/booking"
)

type IBookingService interface {
	NewOutmember(ctx context.Context, agr model.Outmember) (*model.Booking, error)
}

type Service struct {
	bookingpb.UnsafeFinalReportServiceServer
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

func NewService(opts Opts) bookingpb.FinalReportServiceServer {
	return &Service{store: opts.Store, logger: opts.Logger, bookingService: opts.BookingService, notificator: opts.Notificator}
}
