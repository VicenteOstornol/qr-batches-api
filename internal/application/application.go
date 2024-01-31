package application

import (
	"fmt"

	"github.com/VicenteOstornol/lotesapi/config"
	"github.com/VicenteOstornol/lotesapi/internal/delivery"
	"github.com/VicenteOstornol/lotesapi/internal/repository"
	batchRepository "github.com/VicenteOstornol/lotesapi/internal/repository/batch"
	userRepository "github.com/VicenteOstornol/lotesapi/internal/repository/user"
	"github.com/VicenteOstornol/lotesapi/internal/service"
	batchService "github.com/VicenteOstornol/lotesapi/internal/service/batch"
	userService "github.com/VicenteOstornol/lotesapi/internal/service/user"
	_ "github.com/VicenteOstornol/lotesapi/pkg/migrations"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type Application struct {
	name      string
	validator *validator.Validate
	config    *config.Config
	delivery  *fiber.App
}

type Config struct {
	Name      string
	Validator *validator.Validate
	Config    *config.Config
	Postgres  *sqlx.DB
}

func New(conf *Config) *Application {
	userRepository := userRepository.New(conf.Postgres)
	batchRepository := batchRepository.New(conf.Postgres)
	repository := repository.New(userRepository, batchRepository)

	userService := userService.New(repository)
	batchService := batchService.New(repository)
	service := service.New(userService, batchService)

	delivery := delivery.New(service, conf.Validator)
	return &Application{
		name:      conf.Name,
		validator: conf.Validator,
		config:    conf.Config,
		delivery:  delivery,
	}
}

func (app *Application) Start() error {
	return app.delivery.Listen(fmt.Sprintf(":%s", app.config.HTTPPort))
}
