package notecreator

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog"

	"be/internal/datastore"
	"be/internal/datastore/address"
)

type SchedulerSettings struct {
	Day     int
	Hour    int
	Minutes int
}

func (s SchedulerSettings) getSchedulerString() string {
	return fmt.Sprintf("%d %d %d * *", s.Minutes, s.Hour, s.Day)
}

type Config struct {
	Sender                string
	URLBase               string
	BotID                 uuid.UUID
	BotToken              string
	IP                    string
	BookingViewFrontRoute string
	RoomViewFrontRoute string
	EnabledNotify         bool
	ReportsChecker        SchedulerSettings
	MissedReportsChecker  SchedulerSettings
}

type NoteCreator struct {
	cfg    Config
	store  datastore.IDatastore
	logger zerolog.Logger

	reportsCheckerScheduler       *cron.Cron
	missedReportsCheckerScheduler *cron.Cron

	getLink func(id, ip, secret string, addressType address.TypeAddress, uuid uuid.UUID) (string, error)
}

func NewService(cfg Config, store datastore.IDatastore, logger zerolog.Logger,
	getLink func(id, ip, secret string, addressType address.TypeAddress, uuid uuid.UUID) (string, error),
) *NoteCreator {
	n := &NoteCreator{
		cfg:     cfg,
		store:   store,
		logger:  logger,
		getLink: getLink,
	}

	return n
}

func (n *NoteCreator) Run() {
	if n.cfg.EnabledNotify {
		n.reportsChecker(context.Background())
		n.missedReportsChecker(context.Background())
	}
}
