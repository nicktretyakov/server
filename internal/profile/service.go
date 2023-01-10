package profile

import (
	"context"
	"fmt"

	"github.com/imroc/req"
	"github.com/rs/zerolog"
)

type IProfile interface {
	FindEmployeeByPortalCode(ctx context.Context, portalCode uint64) (*Employee, error)
	FindEmployees(ctx context.Context, query string, limit, offset uint64) (EmployeeList, error)
}

func New(apiURL, apiToken string, opts ...ServiceOption) IProfile {
	s := service{
		r:        req.New(),
		apiURL:   apiURL,
		apiToken: apiToken,
	}

	for _, opt := range opts {
		opt(&s)
	}

	return s
}

type service struct {
	r        *req.Req
	apiURL   string
	apiToken string
	logger   zerolog.Logger
}

func (s *service) get(ctx context.Context, path string) (*req.Resp, error) {
	return s.r.Get(s.url(path), ctx, s.headers())
}

func (s *service) headers() req.Header {
	return req.Header{
		"authorization": fmt.Sprintf("bearer %s", s.apiToken),
	}
}

func (s *service) url(path string) string {
	return s.apiURL + path
}
