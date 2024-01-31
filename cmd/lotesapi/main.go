package main

import (
	"log"

	"github.com/VicenteOstornol/lotesapi/config"
	"github.com/VicenteOstornol/lotesapi/internal/application"
	"github.com/VicenteOstornol/lotesapi/internal/repository/postgres"
	"github.com/go-playground/validator/v10"
)

func main() {
	conf, err := config.New()
	if err != nil {
		log.Fatalf("Main: config.New error: %v", err)
	}

	db, err := postgres.New(&conf.Postgres)
	if err != nil {
		log.Fatalf("Main: postgres.New error: %v", err)
	}

	validate := validator.New()
	app := application.New(&application.Config{
		Name:      "lotesapi-wrapper",
		Validator: validate,
		Config:    conf,
		Postgres:  db,
	})

	err = app.Start()
	if err != nil {
		log.Fatalf("Main: app.Start error: %v", err)
	}
}
