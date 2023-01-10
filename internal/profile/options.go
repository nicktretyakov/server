package profile

import (
	"net/http"

	"github.com/rs/zerolog"
)

type ServiceOption func(s *service)

func WithClient(h *http.Client) ServiceOption {
	return func(s *service) {
		s.r.SetClient(h)
	}
}

func WithLogger(l zerolog.Logger) ServiceOption {
	return func(s *service) {
		s.logger = l
	}
}
