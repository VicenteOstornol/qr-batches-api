package migrations

import (
	"fmt"
	"log"

	"github.com/VicenteOstornol/lotesapi/config"
	_ "github.com/VicenteOstornol/lotesapi/internal/repository/postgres/migrations"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func init() {
	conf, err := config.New()
	if err != nil {
		log.Fatalf("migrations: init config.New error: %v", err)
	}

	psURI := fmt.Sprintf(
		"host=%s dbname=%s user=%s port=%s password=%s sslmode=disable",
		conf.Host, conf.DBName, conf.Username, conf.Port, conf.Password,
	)

	db, err := goose.OpenDBWithDriver("postgres", psURI)
	if err != nil {
		log.Fatalf("migrations: init goose.OpenDBWithDriver error: %v", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: init db.Close error: %v", err)
		}
	}()

	if err := goose.Up(db, "./internal/repository/postgres/migrations"); err != nil {
		log.Fatalf("migrations: init goose.Up error: %v", err)
	}
}
