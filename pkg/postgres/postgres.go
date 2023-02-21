package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

const (
	maxConn           = 10
	minConns          = 5
	healthCheckPeriod = 3 * time.Minute
	maxConnIdleTime   = 1 * time.Minute
	maxConnLifetime   = 3 * time.Minute
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DB       string
	SSLMode  string
}

func NewConnPool(config *Config) (*pgxpool.Pool, error) {
	connString := buildConnString(config)
	poolCfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, errors.WithMessage(err, "pgxpool parse config")
	}

	poolCfg.MaxConns = maxConn
	poolCfg.MinConns = minConns
	poolCfg.HealthCheckPeriod = healthCheckPeriod
	poolCfg.MaxConnIdleTime = maxConnIdleTime
	poolCfg.MaxConnLifetime = maxConnLifetime

	ctx := context.Background()
	connPool, err := pgxpool.ConnectConfig(ctx, poolCfg)
	if err != nil {
		return nil, errors.WithMessage(err, "pgxpool new with config")
	}

	return connPool, nil
}

func buildConnString(config *Config) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DB,
		config.SSLMode,
	)
}
