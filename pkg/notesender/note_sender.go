package notesender

import (
	"context"
	"time"

	"be/internal/bot"
	"be/internal/notes/emailsender"

	"github.com/rs/zerolog"

	"be/internal/datastore"
)

type Config struct {
	EmailPeriod time.Duration
	LifePeriod  time.Duration
	BotID       string
	BotToken    string
	IP          string
}

const (
	Private = 0
	Channel = 2
)

type NoteSender struct {
	cfg         Config
	store       datastore.IDatastore
	logger      zerolog.Logger
	emailSender emailsender.IEmailSender
	chatSender  bot.ChatAPI
}

func NewService(
	cfg Config,
	logger zerolog.Logger,
	store datastore.IDatastore,
	emailSender emailsender.IEmailSender,
	chatSender bot.ChatAPI,
) *NoteSender {
	n := &NoteSender{
		cfg:         cfg,
		store:       store,
		logger:      logger,
		emailSender: emailSender,
		chatSender:  chatSender,
	}

	return n
}

func (s *NoteSender) Run() {
	s.emailWorker(context.Background())
	s.lifeWorker(context.Background())
}
