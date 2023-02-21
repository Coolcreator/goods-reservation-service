package migration

import (
	"embed"
	"strings"

	"github.com/pkg/errors"

	_ "github.com/golang-migrate/migrate/v4/database/pgx"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed sql
var migrations embed.FS

func Migrate(URI string) error {
	source, err := iofs.New(migrations, "sql")
	if err != nil {
		return errors.WithMessage(err, "iofs new")
	}

	m, err := migrate.NewWithSourceInstance("iofs", source, strings.Replace(URI, "postgres://", "pgx://", 1))
	if err != nil {
		return errors.WithMessage(err, "new with source instance")
	}
	defer m.Close()

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return errors.WithMessage(err, "migrate up")
	}

	return nil
}
