package delivery

import (
	"github.com/VicenteOstornol/lotesapi/internal/delivery/batch"
	"github.com/VicenteOstornol/lotesapi/internal/delivery/user"
	"github.com/VicenteOstornol/lotesapi/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func New(service *service.Service, validate *validator.Validate) *fiber.App {
	fiberApp := fiber.New()
	fiberApp.Use(cors.New())

	api := fiberApp.Group("/api")
	user.New(service, validate, api)
	batch.New(service, validate, api)
	return fiberApp
}
