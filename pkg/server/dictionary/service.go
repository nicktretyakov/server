package dictionary

import (
	"github.com/rs/zerolog"
	bookingpb "be/proto"

	"be/internal/datastore"
)

type Service struct {
	bookingpb.UnsafeDictionaryServiceServer
	store  datastore.IDatastore
	logger zerolog.Logger
}

func NewService(store datastore.IDatastore, logger zerolog.Logger) bookingpb.DictionaryServiceServer {
	return &Service{store: store, logger: logger}
}
