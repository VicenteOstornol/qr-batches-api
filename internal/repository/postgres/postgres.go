package postgres

import (
	"fmt"

	"github.com/VicenteOstornol/lotesapi/config"
	"github.com/VicenteOstornol/lotesapi/pkg/errors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func New(conf *config.Postgres) (*sqlx.DB, error) {
	psURI := fmt.Sprintf(
		"host=%s dbname=%s user=%s port=%s password=%s sslmode=disable",
		conf.Host, conf.DBName, conf.Username, conf.Port, conf.Password,
	)

	db, err := sqlx.Connect("postgres", psURI)
	if err != nil {
		return nil, errors.Wrap(err, "postgres: New sql.Open error")
	}

	if err := db.Ping(); err != nil {
		return nil, errors.Wrap(err, "postgres: New db.Ping error")
	}

	return db, nil
}
