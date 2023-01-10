package base

import (
	"context"
	"fmt"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/zerologadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog"
)

type Opts struct {
	DSN    string
	Logger zerolog.Logger
}

type DB struct {
	Conn   *pgxpool.Pool
	logger zerolog.Logger
}

func New(opts Opts) (*DB, error) {
	var (
		cfg *pgxpool.Config
		err error
	)

	if cfg, err = pgxpool.ParseConfig(opts.DSN); err != nil {
		return nil, fmt.Errorf("unable to parse DATABASE_URL: %w", err)
	}

	if opts.Logger.GetLevel() == zerolog.DebugLevel {
		cfg.ConnConfig.Logger = zerologadapter.NewLogger(opts.Logger)
	}

	conn, err := pgxpool.ConnectConfig(context.Background(), cfg)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	return &DB{
		logger: opts.Logger,
		Conn:   conn,
	}, nil
}

func (db *DB) ExecBuilder(ctx context.Context, builder sq.Sqlizer) (pgconn.CommandTag, error) {
	rawSQL, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var cmd pgconn.CommandTag

	if cmd, err = db.Conn.Exec(ctx, rawSQL, args...); err != nil {
		return nil, recoverDBError(err)
	}

	return cmd, nil
}

func (db *DB) ExecTxBuilder(ctx context.Context, tx pgx.Tx, insertBuilder sq.Sqlizer) (pgconn.CommandTag, error) {
	rawSQL, args, err := insertBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var cmd pgconn.CommandTag

	if cmd, err = tx.Exec(ctx, rawSQL, args...); err != nil {
		return nil, recoverDBError(err)
	}

	return cmd, nil
}

func (db *DB) Get(ctx context.Context, sqlBuilder sq.Sqlizer, dst interface{}) error {
	rawSQL, args, err := sqlBuilder.ToSql()
	if err != nil {
		return err
	}

	return recoverDBError(pgxscan.Get(ctx, db, dst, rawSQL, args...))
}

func (db *DB) GetTx(ctx context.Context, tx pgx.Tx, sqlBuilder sq.Sqlizer, dst interface{}) error {
	rawSQL, args, err := sqlBuilder.ToSql()
	if err != nil {
		return err
	}

	return recoverDBError(pgxscan.Get(ctx, tx, dst, rawSQL, args...))
}

func (db *DB) Select(ctx context.Context, sqlBuilder sq.SelectBuilder, dst interface{}) error {
	rawSQL, args, err := sqlBuilder.ToSql()
	if err != nil {
		return err
	}

	log.Println("sql ", rawSQL)
	log.Println("args ", args)

	return recoverDBError(pgxscan.Select(ctx, db, dst, rawSQL, args...))
}

func (db *DB) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	return db.Conn.Query(ctx, query, args...)
}

func (db *DB) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return db.Conn.Exec(ctx, query, args...)
}

func (db *DB) Close() {
	db.Conn.Close()
}

func (db *DB) Config() *pgxpool.Config {
	return db.Conn.Config()
}

// Now return postgres timestamptz with Resolution = 1 microsecond
// https://postgrespro.ru/docs/postgrespro/12/datatype-datetime
func (db *DB) Now() time.Time {
	return time.Now().In(time.UTC).Round(time.Microsecond)
}
