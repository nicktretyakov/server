package room

import (
	"github.com/rs/zerolog"
	bookingpb "be/proto"

	"be/internal/datastore"
	"be/internal/notecreator"
	"be/internal/room"
)

type Config struct{}

type Service struct {
	bookingpb.UnsafeRoomServiceServer
	store          datastore.IDatastore
	roomService room.IRoomService
	logger         zerolog.Logger
	notificator    *notecreator.NoteCreator
}

func NewService(store datastore.IDatastore,
	roomService room.IRoomService,
	notificator *notecreator.NoteCreator,
	opts ...Option,
) bookingpb.RoomServiceServer {
	s := &Service{
		store:          store,
		logger:         zerolog.Logger{},
		roomService: roomService,
		notificator:    notificator,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

type Option func(s *Service)

func WithLogger(logger zerolog.Logger) Option {
	return func(s *Service) {
		s.logger = logger
	}
}
