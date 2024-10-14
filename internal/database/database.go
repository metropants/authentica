package database

import (
	"context"
	"database/sql"
	"embed"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/metropants/authentica/internal/config"
	"github.com/pressly/goose/v3"
)

const (
	Dialect            = "postgres"
	Driver             = "pgx/v5"
	MigrationDirectory = "migrations"
)

//go:embed migrations/*.sql
var migrations embed.FS

func New(c *config.Config) (*pgxpool.Pool, error) {
	uri := c.DatabaseURI
	if len(uri) <= 0 {
		return nil, errors.New("invalid database connection string")
	}

	pool, err := pgxpool.New(context.Background(), uri)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(context.TODO())
	if err != nil {
		return nil, errors.New("unable to ping database connection")
	}

	return pool, nil
}

func Migrate(pool *pgxpool.Pool) error {
	if pool == nil {
		return errors.New("database pool is nil")
	}

	goose.SetBaseFS(migrations)
	if err := goose.SetDialect(Dialect); err != nil {
		return err
	}

	conn := pool.Config().ConnConfig.ConnString()
	db, err := sql.Open(Driver, conn)
	if err != nil {
		return err
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	if err := goose.Up(db, MigrationDirectory); err != nil {
		return err
	}
	return nil
}
