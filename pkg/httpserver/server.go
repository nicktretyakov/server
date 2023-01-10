package httpserver

import (
	"context"
	"errors"
	"io"
	"net"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

type Config struct {
	IP   string
	Port string
}

func (c Config) addr() string {
	return net.JoinHostPort(c.IP, c.Port)
}

type Server struct {
	cfg        Config
	log        zerolog.Logger
	rootRouter *mux.Router

	srv *http.Server
}

//nolint:gosec
func New(cfg Config, log zerolog.Logger) *Server {
	rootRouter := mux.NewRouter()

	return &Server{
		cfg:        cfg,
		log:        log,
		rootRouter: rootRouter,
		srv: &http.Server{
			Addr:    cfg.addr(),
			Handler: handlers.CustomLoggingHandler(log, handlers.CORS()(rootRouter), getHandlerLogger(log)),
		},
	}
}

func (s *Server) Run() error {
	go func() {
		s.log.Info().Msgf("Running server on %s", s.cfg.addr())

		if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.log.Fatal().Err(err).Msgf("ListenAndServe fail")
		}
	}()

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func (s *Server) GetRouter() *mux.Router {
	return s.rootRouter
}

func getHandlerLogger(log zerolog.Logger) handlers.LogFormatter {
	return func(_ io.Writer, params handlers.LogFormatterParams) {
		event := log.Info()
		if params.StatusCode >= http.StatusBadRequest {
			event = log.Error()
		}

		log.Debug().
			Fields(map[string]interface{}{
				"query":   params.URL.Query(),
				"headers": params.Request.Header,
			}).
			Msgf("%s %s", params.Request.Method, params.Request.URL.String())

		event.Int("status", params.StatusCode).
			Time("timestamp", params.TimeStamp.UTC()).
			Str("url", params.URL.String()).
			Msgf("%s %s", params.Request.Method, params.Request.URL.String())
	}
}
