package testutil

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path"
	"strings"
	"testing"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"

	// driver.
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/rs/xid"

	"be/internal/datastore"
	"be/internal/datastore/base"
	"be/pkg/logging"
)

type DBResource struct {
	DB       *base.DB
	Fixtures *FixtureStore
}

type TestDBConfig struct {
	PostgresDsn           string `env:"TN_POSTGRES_DSN,required"`
	Migrations            string `env:"TN_MIGRATION_MIGRATIONS,required" envDefault:""`
	MigrationVersionTable string `env:"TN_MIGRATION_VERSION_TABLE" envDefault:"public.schema_version"`
	LogLevel              string `env:"TN_LOGLEVEL" envDefault:"warn"`
}

func SetupTestDataBase(t *testing.T) (resource *DBResource, err error) { //nolint:funlen
	t.Helper()

	var (
		cfg TestDBConfig
		ctx = context.Background()
	)

	if err = env.Parse(&cfg); err != nil {
		return nil, err
	}

	dsnConfig, err := pgconn.ParseConfig(cfg.PostgresDsn)
	if err != nil {
		return nil, err
	}

	testDataBaseName := "test_" + xid.New().String() + "_" + dsnConfig.Database
	newDsn := fmt.Sprintf("user=%s password=%s host=%s port=%v database=%s sslmode=disable",
		dsnConfig.User,
		dsnConfig.Password,
		dsnConfig.Host,
		dsnConfig.Port,
		testDataBaseName)

	sqlConn, err := pgx.Connect(ctx, cfg.PostgresDsn)
	if err != nil {
		return nil, err
	}

	err = dropDataBase(ctx, sqlConn, testDataBaseName)
	if err != nil {
		return nil, err
	}

	_, err = sqlConn.Exec(ctx, "create database "+testDataBaseName)
	if err != nil {
		return nil, err
	}

	cfg.PostgresDsn = newDsn

	if err = datastore.AutoMigrate(datastore.DBConf{
		MigrationsPath: cfg.Migrations,
		Version:        "last",
		Auto:           true,
		VersionTable:   cfg.MigrationVersionTable,
		DSN:            cfg.PostgresDsn,
		LogLevel:       cfg.LogLevel,
	}, true); err != nil {
		return nil, err
	}

	logger := logging.GetLogger(logging.ConsoleLoggerType, cfg.LogLevel)

	db, err := base.New(base.Opts{
		DSN:    cfg.PostgresDsn,
		Logger: logger,
	})
	if err != nil {
		return nil, err
	}

	tearDown := func() {
		db.Close()

		dropErr := dropDataBase(ctx, sqlConn, testDataBaseName)
		if dropErr != nil {
			logger.Err(dropErr).Msgf("cannot drop database %s", testDataBaseName)
		}
	}

	t.Cleanup(tearDown)

	return &DBResource{
		DB: db,
	}, nil
}

// SetupTestDataBaseWithFixtures creates test database.
// Registers teardown function.
func SetupTestDataBaseWithFixtures(t *testing.T) (resource *DBResource, err error) {
	t.Helper()

	ds, err := SetupTestDataBase(t)
	if err != nil {
		return nil, err
	}

	err = LoadFixtures(ds.DB.Config().ConnConfig.ConnString(), nil)
	if err != nil {
		return nil, err
	}

	ds.Fixtures = NewFixtureStore(t)

	return ds, nil
}

func dropDataBase(ctx context.Context, sqlConn *pgx.Conn, dbName string) (err error) {
	_, err = sqlConn.Exec(ctx, "SELECT pg_terminate_backend(pid) FROM pg_stat_activity WHERE datname = $1", dbName)
	if err != nil {
		return err
	}

	_, err = sqlConn.Exec(ctx, "drop database if exists "+dbName)
	if err != nil {
		return err
	}

	return nil
}

func LoadFixtures(pgConn string, _dirPath *string) error {
	dirPath := "../testutil/fixtures"
	if _dirPath != nil {
		dirPath = *_dirPath
	}

	fixtures, err := fixturesList(dirPath)
	if err != nil {
		return err
	}

	db, err := sql.Open("pgx", pgConn)
	if err != nil {
		return err
	}

	fixturesLoader, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("pgx"),
		testfixtures.Files(fixtures...),
		testfixtures.Location(time.UTC),
		testfixtures.DangerousSkipTestDatabaseCheck(),
	)
	if err != nil {
		return err
	}

	return fixturesLoader.Load()
}

func fixturesList(dirPath string) ([]string, error) {
	filesPath := make([]string, 0)

	files, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		if strings.Contains(f.Name(), ".yml") {
			filesPath = append(filesPath, path.Join(dirPath, f.Name()))
		}
	}

	return filesPath, nil
}
