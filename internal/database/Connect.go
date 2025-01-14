package database

import (
	"context"
	"sync"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	dbPool *pgxpool.Pool
	once sync.Once
	dbErr error
)

func Connect() (*pgxpool.Pool, error) {
	once.Do(func() {
		ctx := context.Background()

		config, err := pgxpool.ParseConfig("postgres://rakhatmelsovforjobgmail.com@127.0.0.1:5432/postdb")
			if err != nil {
			dbErr = err
		}
		dbPool, err = pgxpool.NewWithConfig(ctx, config)
		if err != nil {
			dbErr = err
		}

	})
	if dbErr != nil {
		return nil, dbErr
	}
	return dbPool, nil
}
