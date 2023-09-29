package database

import (
	"context"
	"github.com/giusepperoro/todolist.git/intenal/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

func New(ctx context.Context, cfg config.ServiceConfiguration) (*pgxpool.Pool, error) {
	connection, err := pgxpool.Connect(ctx, cfg.PostgresConnectUrl)
	if err != nil {
		return nil, err
	}
	err = connection.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return connection, nil
}
