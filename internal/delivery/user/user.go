package user

import (
	"net/http"

	"github.com/VicenteOstornol/lotesapi/entities"
	"github.com/VicenteOstornol/lotesapi/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserDelivery struct {
	service   *service.Service
	validator *validator.Validate
}

func New(service *service.Service, validate *validator.Validate, router fiber.Router) {
	delivery := &UserDelivery{service: service, validator: validate}
	route := router.Group("/user")
	delivery.CreateUser(route)
	delivery.GetBatchesByUserID(route)
}

func (p *UserDelivery) CreateUser(route fiber.Router) {
	route.Post("/create", func(c *fiber.Ctx) error {
		var user entities.User
		if err := c.BodyParser(&user); err != nil {
			return fiber.NewError(http.StatusUnprocessableEntity, err.Error())
		}

		if err := p.validator.StructCtx(c.Context(), &user); err != nil {
			return fiber.NewError(http.StatusBadRequest, err.Error())
		}

		createdUser, err := p.service.User.Create(c.Context(), &user)
		if err != nil {
			return fiber.NewError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(fiber.Map{
			"data": createdUser,
		})
	})
}

func (p *UserDelivery) GetBatchesByUserID(route fiber.Router) {
	route.Get("/:id/batches", func(c *fiber.Ctx) error {
		id := c.Params("id")
		batches, err := p.service.User.GetBatchesByUserID(c.Context(), id)
		if err != nil {
			return fiber.NewError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(fiber.Map{
			"data": batches,
		})
	})
}
