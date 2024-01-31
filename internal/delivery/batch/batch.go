package batch

import (
	"net/http"

	"github.com/VicenteOstornol/lotesapi/entities"
	"github.com/VicenteOstornol/lotesapi/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type BatchDelivery struct {
	service   *service.Service
	validator *validator.Validate
}

func New(service *service.Service, validate *validator.Validate, router fiber.Router) {
	delivery := &BatchDelivery{service: service, validator: validate}
	route := router.Group("/batch")
	delivery.CreateBatch(route)
	delivery.DownloadPDFBatch(route)
	delivery.GetBatchWithQR(route)
}

func (p *BatchDelivery) CreateBatch(route fiber.Router) {
	route.Post("/create", func(c *fiber.Ctx) error {
		var batch entities.Batch
		if err := c.BodyParser(&batch); err != nil {
			return fiber.NewError(http.StatusUnprocessableEntity, err.Error())
		}

		if err := p.validator.StructCtx(c.Context(), &batch); err != nil {
			return fiber.NewError(http.StatusBadRequest, err.Error())
		}

		createdBatch, err := p.service.Batch.Create(c.Context(), &batch)
		if err != nil {
			return fiber.NewError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(fiber.Map{
			"data": createdBatch,
		})
	})
}

func (p *BatchDelivery) DownloadPDFBatch(route fiber.Router) {
	route.Get("/:id/download", func(c *fiber.Ctx) error {
		id := c.Params("id")
		pdf, err := p.service.Batch.DownloadPDF(c.Context(), id)
		if err != nil {
			return fiber.NewError(http.StatusInternalServerError, err.Error())
		}

		c.Response().Header.Set("Content-Type", "application/pdf")
		pdf.Output(c.Response().BodyWriter())
		return c.SendStatus(http.StatusOK)
	})
}

func (p *BatchDelivery) GetBatchWithQR(route fiber.Router) {
	route.Get("/:id/:qr_number", func(c *fiber.Ctx) error {
		id, qr_number := c.Params("id"), c.Params("qr_number")
		batch, err := p.service.Batch.Get(c.Context(), id)
		if err != nil {
			return fiber.NewError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(fiber.Map{
			"data": fiber.Map{
				"batch":     batch,
				"qr_number": qr_number,
			}})
	})
}
