package emailsender

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

type service struct {
	dialer *gomail.Dialer

	// logger zerolog.Logger
}

type Config struct {
	SkipTLS  bool
	Host     string
	Port     int
	User     string
	Password string
}

//nolint:gosec
func New(cfg Config) IEmailSender {
	dialer := gomail.NewDialer(cfg.Host, cfg.Port, cfg.User, cfg.Password)
	if cfg.SkipTLS {
		dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}

	return &service{
		dialer: dialer,
	}
}
