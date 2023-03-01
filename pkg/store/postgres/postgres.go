package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"strings"
	"time"
)

// Store is a system store
type Store struct {
	Pool *pgxpool.Pool
}

// DBConfig is a database config.
type DBConfig struct {
	Host     string
	Database string
	User     string
	Password string
	SSLMode  string
	Port     uint16
}

// toDSN creates a connection string
func (c *DBConfig) toDSN() string {
	var args []string

	if len(c.Host) > 0 {
		args = append(args, fmt.Sprintf("host=%s", c.Host))
	}

	if c.Port > 0 {
		args = append(args, fmt.Sprintf("port=%d", c.Port))
	}

	if len(c.Database) > 0 {
		args = append(args, fmt.Sprintf("dbname=%s", c.Database))
	}

	if len(c.User) > 0 {
		args = append(args, fmt.Sprintf("user=%s", c.User))
	}

	if len(c.Password) > 0 {
		args = append(args, fmt.Sprintf("password=%s", c.Password))
	}

	if len(c.SSLMode) > 0 {
		args = append(args, fmt.Sprintf("sslmode=%s", c.SSLMode))
	}

	return strings.Join(args, " ")
}

// NewPostgres return a new Store
func NewPostgres(cfg DBConfig) (*Store, error) {
	config, err := pgxpool.ParseConfig(cfg.toDSN())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	conn, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err = conn.Ping(ctx); err != nil {
		return nil, errors.WithStack(err)
	}

	return &Store{Pool: conn}, nil
}
