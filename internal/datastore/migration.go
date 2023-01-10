package datastore

import (
	//"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	//"be/tern"
	//"github.com/jackc/tern/migrate"
)

var ErrMigration = errors.New("migration error")

type DBConf struct {
	MigrationsPath string
	Version        string
	Auto           bool
	VersionTable   string
	DSN            string
	LogLevel       string
}

func AutoMigrate(cfg DBConf, disableLogs bool) error {
	if disableLogs {
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	}

	// m, err := tern.New(
	// 	cfg.DSN,
	// 	"",
	// 	cfg.MigrationsPath,
	// 	cfg.VersionTable,
	// 	cfg.Version,
	// 	cfg.LogLevel,
	// 	&map[string]interface{}{},
	// )
	// if err != nil {
	// 	return fmt.Errorf("%w: %s", ErrMigration, err.Error())
	// }

	// if err := m.Run(); err != nil {
	// 	return fmt.Errorf("%w: %s", ErrMigration, err.Error())
	// }

	return nil
}
