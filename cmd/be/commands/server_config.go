package commands

import (
	"time"

	"github.com/google/uuid"

	"golang.org/x/oauth2"

	"be/internal/datastore"
	"be/pkg/auth"
	"be/pkg/httpserver"
	grpcServer "be/pkg/server/server"
)

type serverConfig struct {
	LogLevel  string `env:"TN_LOGLEVEL" envDefault:"warn"`   // debug, info, warn, error, fatal, ""
	LogOutput string `env:"TN_LOG_OUTPUT" envDefault:"json"` // json, console, ""
	IP        string `env:"TN_IP" envDefault:"0.0.0.0"`
	GRPCPort  string `env:"TN_GRPC_PORT" envDefault:"8080"`
	HTTPPort  string `env:"TN_HTTP_PORT" envDefault:"3015"`

	// Postgres
	PostgresDsn string `env:"TN_POSTGRES_DSN" envDefault:"user=postgres password=nick host=localhost port=5432 database=postgres sslmode=disable"`
	// Postgres end

	//	Profile
	ProfileAPIURL   string `env:"TN_PROFILE_EMPLOYEE_API_URL" envDefault:"https://www.tn-profile-accept.tjdev.ru"`
	ProfileAPIToken string `env:"TN_PROFILE_EMPLOYEE_API_TOKEN" envDefault:"wvrabvlttxhnr0l0tvrrmk5pmdbnv1f6tfdjne0ywxrnr05rt0drefpqrxlnvgxtq2c9pqo"`

	ClientID     string   `env:"TN_PROFILE_CLIENT_ID" envDefault:"tn-booking-mini-app"`
	ClientSecret string   `env:"TN_PROFILE_CLIENT_SECRET" envDefault:"TnpFMVpXWmhZakV0TXpVM09TMDBaREl4TFdJNE5HWXRNV0V3WmpFMk1qUmhNRGcwQ2c9PQ"`
	AuthURL      string   `env:"TN_PROFILE_AUTH_URL" envDefault:"https://www.tn-profile-accept.tjdev.ru/api/v1/oauth/authorize"`
	TokenURL     string   `env:"TN_PROFILE_TOKEN_URL" envDefault:"https://www.tn-profile-accept.tjdev.ru/api/v1/oauth/token"`
	RedirectURL  string   `env:"TN_PROFILE_REDIRECT_URL" envDefault:"https://miniapp-tn-booking-dev.tages.dev"`
	Scopes       []string `env:"TN_PROFILE_SCOPES" envSeparator:"," envDefault:"phone.read,employee.read,email.read"`
	// Profile end

	// Sessions
	AccessExpiry  time.Duration `env:"TN_ACCESS_EXPIRY" envDefault:"1h"`
	RefreshExpiry time.Duration `env:"TN_REFRESH_EXPIRY" envDefault:"48h"`
	Secret        string        `env:"TN_SECRET" envDefault:"oa05hvmPNa2lv0WLRrEu55MDbpRee4xtf5rnm05tTX1rave0yWWpgALlUQ2tfn5vek0q2c"`
	// Sessions end

	// email sender
	SkipTLS       bool   `env:"TN_SKIP_TLS" envDefault:"false"`
	EmailHost     string `env:"TN_EMAIL_HOST" envDefault:"smtp.gmail.com"`
	EmailPort     int    `env:"TN_EMAIL_PORT" envDefault:"587"`
	EmailUser     string `env:"TN_EMAIL_USER" envDefault:"tn.booking.app@gmail.com"`
	EmailPassword string `env:"TN_EMAIL_PASSWORD" envDefault:"TnPass#10"`
	EmailFrom     string `env:"TN_EMAIL_FROM" envDefault:"tn.booking.app@gmail.com"`
	// email sender end

	// Notes
	ReportsCheckerSchedulerDay     int `env:"TN_REPORTS_CHECKER_SCHEDULER_DAY" envDefault:"8"`
	ReportsCheckerSchedulerHour    int `env:"TN_REPORTS_CHECKER_SCHEDULER_HOUR" envDefault:"07"`
	ReportsCheckerSchedulerMinutes int `env:"TN_REPORTS_CHECKER_SCHEDULER_MINUTES" envDefault:"30"`

	MissedReportsCheckerSchedulerDay     int `env:"TN_MISSED_REPORTS_CHECKER_SCHEDULER_DAY" envDefault:"10"`
	MissedReportsCheckerSchedulerHour    int `env:"TN_MISSED_REPORTS_CHECKER_SCHEDULER_HOUR" envDefault:"07"`
	MissedReportsCheckerSchedulerMinutes int `env:"TN_MISSED_REPORTS_CHECKER_SCHEDULER_MINUTES" envDefault:"30"`

	EnableCreateNotes    bool          `env:"TN_ENABLE_CREATE_NOTES" envDefault:"true"`
	SendEmailNotesPeriod time.Duration `env:"TN_SEND_EMAIL_NOTES_PERIOD" envDefault:"1m"`
	SendLifeNotesPeriod  time.Duration `env:"TN_SEND_LIFE_NOTES_PERIOD" envDefault:"1m"`

	URLBaseNote               string    `env:"TN_URL_BASE_NOTE" envDefault:"https://tn-office-dev.tages.dev"`
	BookingViewFrontRouteNote string    `env:"TN_BOOKING_VIEW_FRONT_ROUTE_NOTE" envDefault:"/booking/[[uuid]]"`
	RoomViewFrontRouteNote string    `env:"TN_ROOM_VIEW_FRONT_ROUTE_NOTE" envDefault:"/room/[[uuid]]"`
	BotURL                            string    `env:"TN_BOT_API_URL" envDefault:"https://bots-tn-life-dev.tages.dev/bots"`
	BotID                             uuid.UUID `env:"TN_BOT_ID" envDefault:"a6c11aaf-b133-4256-8c9c-61afa71492a6"`
	BotToken                          string    `env:"TN_BOT_TOKEN" envDefault:"Cfuy6LwptrKJgRM7u9kX7l4Bi3EZzyZKmgz4zXGD"`
	// Notes end
}

func (c serverConfig) dbConf() datastore.DBConf {
	return datastore.DBConf{
		//MigrationsPath: c.Migrations,
		Version:        "last",
		Auto:           true,
		//VersionTable:   c.MigrationVersionTable,
		DSN:            c.PostgresDsn,
		LogLevel:       c.LogLevel,
	}
}

func (c serverConfig) authConfig() auth.Config {
	return auth.Config{
		OauthCfg: oauth2.Config{
			ClientID:     c.ClientID,
			ClientSecret: c.ClientSecret,
			Endpoint: oauth2.Endpoint{
				AuthURL:  c.AuthURL,
				TokenURL: c.TokenURL,
			},
			RedirectURL: c.RedirectURL,
			Scopes:      c.Scopes,
		},
		AccessExpiry:  c.AccessExpiry,
		RefreshExpiry: c.RefreshExpiry,
		Secret:        []byte(c.Secret),
	}
}

func (c serverConfig) grpcServerConfig() grpcServer.Config {
	return grpcServer.Config{
		Secret: c.Secret,
	}
}

func (c serverConfig) httpServerConfig() httpserver.Config {
	return httpserver.Config{
		IP:   c.IP,
		Port: c.HTTPPort,
	}
}


