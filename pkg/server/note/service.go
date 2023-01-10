package note

import (
	"github.com/rs/zerolog"
	bookingpb "be/proto"

	"be/internal/datastore"
)

type Config struct{}

type Service struct {
	bookingpb.UnsafeNoteServer
	store  datastore.IDatastore
	logger zerolog.Logger
}

func NewService(store datastore.IDatastore, opts ...Option) bookingpb.NoteServer {
	s := &Service{
		store:  store,
		logger: zerolog.Logger{},
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
