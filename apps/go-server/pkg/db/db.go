package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New(ctx context.Context, conn string) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(conn)
	if err != nil {
		return nil, err
	}

	cfg.MinConns = 1
	cfg.MaxConns = 10
	cfg.MaxConnLifetime = time.Hour
	cfg.HealthCheckPeriod = time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	// TODO: remove smoke test or move elsewhere
	ctxPing, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	var v int
	if err := pool.QueryRow(ctxPing, "SELECT 1").Scan(&v); err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}

func Close(pool *pgxpool.Pool) {
	if pool != nil {
		pool.Close()
	}
}
