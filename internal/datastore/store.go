package datastore

import (
	"github.com/rs/zerolog"

	outmemberStore "be/internal/datastore/outmember"
	"be/internal/datastore/base"
	dictionaryStore "be/internal/datastore/dictionary"
	finalReportStore "be/internal/datastore/finalreports"
	addressStore "be/internal/datastore/address"
	roomStore "be/internal/datastore/room"
	bookingStore "be/internal/datastore/booking"
	reportStore "be/internal/datastore/report"
	sessionStore "be/internal/datastore/session"
	stageStorage "be/internal/datastore/stage"
	userStorage "be/internal/datastore/user"
	"be/internal/filestorage"
)

type Opts struct {
	DSN         string
	Logger      zerolog.Logger
	FileStorage filestorage.IFileStorage
}

func New(opts Opts) (IDatastore, error) {
	db, err := base.New(base.Opts{
		DSN:    opts.DSN,
		Logger: opts.Logger,
	})
	if err != nil {
		return nil, err
	}

	return &store{db: db, fileStorage: opts.FileStorage}, nil
}

type store struct {
	db          *base.DB
	fileStorage filestorage.IFileStorage
}

func (s store) FinalReport() IFinalReportStore {
	return finalReportStore.New(s.db)
}

func (s store) Outmember() IOutmemberStore {
	return outmemberStore.New(s.db)
}

func (s store) Session() ISessionStore {
	return sessionStore.New(s.db)
}

func (s store) User() IUserStore {
	return userStorage.New(s.db)
}

func (s store) Booking() IBookingStore {
	return bookingStore.New(s.db, s.fileStorage)
}

func (s store) Report() IReportStore {
	return reportStore.New(s.db)
}

func (s store) Dictionary() IDictionaryStore {
	return dictionaryStore.New(s.db)
}

func (s store) Stage() IStageStore {
	return stageStorage.New(s.db, s.fileStorage)
}

func (s store) Address() IAddressStore {
	return addressStore.New(s.db, s.fileStorage)
}

func (s store) Room() IRoomStore {
	return roomStore.New(s.db, s.fileStorage)
}
