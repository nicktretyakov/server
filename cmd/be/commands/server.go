package commands

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"

	"be/internal/outmember"
	"be/internal/bot"
	"be/internal/datastore"
	//"be/internal/filestorage"
	"be/internal/address"
	"be/internal/notecreator"
	"be/internal/notes/emailsender"
	"be/internal/room"
	"be/internal/profile"
	"be/internal/booking"
	"be/internal/report"
	"be/internal/user"
	"be/pkg/auth"
	"be/pkg/handlers"
	"be/pkg/httpserver"
	"be/pkg/logging"
	"be/pkg/notesender"
	grpcServer "be/pkg/server/server"
	"be/pkg/version"
)

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "run",
		Short: "Runs grpcServer",
		RunE:  runServer,
	})
}

//nolint:funlen
func runServer(c *cobra.Command, _ []string) error {
	cfg, err := getConfig()
	if err != nil {
		return err
	}

	logger := logging.GetLogger(cfg.LogOutput, cfg.LogLevel)

	logger.Info().Msgf("Starting server...")
	logger.Info().Msgf("Version: %s", version.Version)

	if err = datastore.AutoMigrate(cfg.dbConf(), false); err != nil {
		return err
	}

	// fs, err := filestorage.New(cfg.Aws())
	// if err != nil {
	// 	return errors.Wrap(err, "s3 failed")
	// }

	storage, err := datastore.New(datastore.Opts{
		DSN:         cfg.PostgresDsn,
		Logger:      logger,
		//FileStorage: fs,
	})
	if err != nil {
		return errors.Wrap(err, "database connection failed")
	}

	profileAPI := profile.New(
		cfg.ProfileAPIURL,
		cfg.ProfileAPIToken,
		profile.WithClient(&http.Client{Timeout: 20 * time.Second}), //nolint:gomnd
		profile.WithLogger(logger),
	)

	chatService := bot.NewChatService(bot.Config{
		BotURL:                            cfg.BotURL,
		BookingViewFrontRouteNote: cfg.BookingViewFrontRouteNote,
		RoomViewFrontRouteNote: cfg.RoomViewFrontRouteNote,
	})

	noteCreator := notecreator.NewService(notecreator.Config{
		Sender:        cfg.EmailFrom,
		EnabledNotify: cfg.EnableCreateNotes,
		ReportsChecker: notecreator.SchedulerSettings{
			Day:     cfg.ReportsCheckerSchedulerDay,
			Hour:    cfg.ReportsCheckerSchedulerHour,
			Minutes: cfg.ReportsCheckerSchedulerMinutes,
		},
		MissedReportsChecker: notecreator.SchedulerSettings{
			Day:     cfg.MissedReportsCheckerSchedulerDay,
			Hour:    cfg.MissedReportsCheckerSchedulerHour,
			Minutes: cfg.MissedReportsCheckerSchedulerMinutes,
		},
		URLBase:               cfg.URLBaseNote,
		BookingViewFrontRoute: cfg.BookingViewFrontRouteNote,
		RoomViewFrontRoute: cfg.RoomViewFrontRouteNote,
		BotID:                 cfg.BotID,
		BotToken:              cfg.BotToken,
		IP:                    cfg.IP,
	},
		storage,
		logger,
		chatService.CreateProposalViewLink,
	)

	noteSender := notesender.NewService(notesender.Config{
		EmailPeriod: cfg.SendEmailNotesPeriod,
		LifePeriod:  cfg.SendLifeNotesPeriod,
		BotID:       cfg.BotID.String(),
		BotToken:    cfg.BotToken,
		IP:          cfg.IP,
	},
		logger,
		storage,
		emailsender.New(emailsender.Config{
			SkipTLS:  cfg.SkipTLS,
			Host:     cfg.EmailHost,
			Port:     cfg.EmailPort,
			User:     cfg.EmailUser,
			Password: cfg.EmailPassword,
		}),
		chatService,
	)

	authApp := auth.New(cfg.authConfig(), toAuthUser(storage, profileAPI), logger)
	srv := grpcServer.New(cfg.grpcServerConfig(), grpcServer.Opts{
		AuthService:         authApp,
		Store:               storage,
		Logger:              logger,
		BookingService:      booking.New(storage, user.New(storage.User(), profileAPI), report.New(storage.Report())),
		RoomService:      room.New(storage, user.New(storage.User(), profileAPI)),
		AddressService: address.New(storage, user.New(storage.User(), profileAPI)),
		OutmemberService:    outmember.New(storage, report.New(storage.Report())),
		ProfileAPI:          profileAPI,
		NoteCreator: noteCreator,
		NoteSender:  noteSender,
	})

	httpServer := httpserver.New(cfg.httpServerConfig(), logger)

	if err = handlers.Bind(authApp, storage.User(), profileAPI, /*fs,*/ httpServer.GetRouter()/*, logger*/); err != nil {
		return err
	}

	return run(c.Context(), *cfg, srv, httpServer, logger)
}

func getConfig() (*serverConfig, error) {
	cfg := serverConfig{}

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func run(ctx context.Context, cfg serverConfig, srv *grpcServer.Server, httpServer *httpserver.Server, logger zerolog.Logger) (err error) {
	if err = grpcServer.Run(srv, cfg.IP, cfg.GRPCPort); err != nil {
		return err
	}

	defer srv.Shutdown()

	if err = httpServer.Run(); err != nil {
		return err
	}

	defer func() {
		if err := httpServer.Shutdown(ctx); err != nil {
			logger.Err(err).Msg("Error during shutdown")
		}
	}()

	listen(logger)

	return nil
}

func listen(logger zerolog.Logger) {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)
	sign := <-quit
	logger.Info().Msgf("Shutting down grpcServer... Reason: %s", sign.String())
}
