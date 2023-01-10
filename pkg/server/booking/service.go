package booking

import (
	"github.com/rs/zerolog"
	bookingpb "be/proto"

	"be/internal/datastore"
	"be/internal/filestorage"
	"be/internal/notecreator"
	"be/internal/booking"
)

type Config struct{}

type Service struct {
	bookingpb.UnsafeBookingServiceServer
	cfg            Config
	store          datastore.IDatastore
	bookingService booking.IBookingService
	logger         zerolog.Logger
	fileLoader     filestorage.IFileLoader
	notificator    *notecreator.NoteCreator
}

func NewService(cfg Config, store datastore.IDatastore,
	bookingService booking.IBookingService, notificator *notecreator.NoteCreator, opts ...Option,
) bookingpb.BookingServiceServer {
	s := &Service{
		cfg:            cfg,
		store:          store,
		bookingService: bookingService,
		notificator:    notificator,
		logger:         zerolog.Logger{},
		fileLoader:     filestorage.FileLoader{},
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

func WithFileLoader(fileLoader filestorage.IFileLoader) Option {
	return func(s *Service) {
		s.fileLoader = fileLoader
	}
}
