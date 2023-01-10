package outmember

import (
	"github.com/rs/zerolog"
	bookingpb "be/proto"

	"be/internal/outmember"
	"be/internal/datastore"
	"be/internal/notecreator"
	"be/internal/room"
	"be/internal/booking"
)

type Service struct {
	bookingpb.UnsafeOutmemberServiceServer
	store            datastore.IDatastore
	logger           zerolog.Logger
	bookingService   booking.IBookingService
	roomService   room.IRoomService
	outmemberService outmember.IOutmemberService
	Notificator      *notecreator.NoteCreator
}

type Opts struct {
	Store            datastore.IDatastore
	BookingService   booking.IBookingService
	RoomService   room.IRoomService
	OutmemberService outmember.IOutmemberService
	Logger           zerolog.Logger
	Notificator      *notecreator.NoteCreator
}

func NewService(opts Opts) bookingpb.OutmemberServiceServer {
	return &Service{
		store:            opts.Store,
		logger:           opts.Logger,
		bookingService:   opts.BookingService,
		roomService:   opts.RoomService,
		outmemberService: opts.OutmemberService,
		Notificator:      opts.Notificator,
	}
}
