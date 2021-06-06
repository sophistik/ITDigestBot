package postgres

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type DB struct {
	Session *sql.DB
	Logger  *zap.Logger
}

func NewDB(logger *zap.Logger, cfg Config) (*DB, error) {
	db, err := sql.Open("postgres", cfg.DNS())
	if err != nil {
		return nil, fmt.Errorf("can't open connection to postgres: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("can't ping db: %w", err)
	}

	db.SetConnMaxLifetime(cfg.MaxConnLifetime)
	db.SetMaxIdleConns(cfg.MaxConns)
	db.SetMaxOpenConns(cfg.MaxConns)

	return &DB{
		Session: db,
		Logger:  logger,
	}, nil
}

func (d *DB) CheckConnection() error {
	var err error

	const maxAttempts = 3

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		if err = d.Session.Ping(); err == nil {
			break
		}

		nextAttemptWait := time.Duration(attempt) * time.Second
		d.Logger.Sugar().Errorf("attempt %d: can't establish a connection with the db, wait for %v: %s",
			attempt,
			nextAttemptWait,
			err,
		)
		time.Sleep(nextAttemptWait)
	}

	return fmt.Errorf("can't connect to db: %w", err)
}

func (d *DB) Close() error {
	if err := d.Session.Close(); err != nil {
		return fmt.Errorf("can't close db: %w", err)
	}

	return nil
}

func (d *DB) ValidateQueries(statements []string) error {
	for i := range statements {
		statement, err := d.Session.Prepare(statements[i])
		if err != nil {
			return fmt.Errorf("can't prepare query %q: %w", statements[i], err)
		}

		if err = statement.Close(); err != nil {
			return fmt.Errorf("can't close statement from query %q: %w", statements[i], err)
		}
	}

	return nil
}
