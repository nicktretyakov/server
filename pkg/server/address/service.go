package address

import (
	"github.com/rs/zerolog"
	bookingpb "be/proto"

	"be/internal/datastore"
	addressStore "be/internal/datastore/address"
	"be/internal/filestorage"
	"be/internal/address"
	"be/internal/notecreator"
)

type Service struct {
	bookingpb.UnsafeAddressServiceServer
	addressService address.IAddressService
	logger              zerolog.Logger
	fileLoader          filestorage.IFileLoader
	notificator         *notecreator.NoteCreator
	store               datastore.IDatastore
}

func NewService(
	addressService address.IAddressService,
	store datastore.IDatastore,
	notificator *notecreator.NoteCreator,
	opts ...Option,
) bookingpb.AddressServiceServer {
	s := &Service{
		addressService: addressService,
		logger:              zerolog.Logger{},
		fileLoader:          filestorage.FileLoader{},
		notificator:         notificator,
		store:               store,
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

func getTypeAddress(typeAddress bookingpb.AddressType) addressStore.TypeAddress {
	switch typeAddress {
	case bookingpb.AddressType_BOOKING:
		return addressStore.BookingAddressType
	case bookingpb.AddressType_ROOM:
		return addressStore.RoomAddressType
	case bookingpb.AddressType_UNKNOWN_OBJECT:
		return addressStore.UnknownAddressType
	default:
		return addressStore.UnknownAddressType
	}
}
