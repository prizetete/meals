package adapters

import (
	"komkrit/core"

	"github.com/gofiber/fiber/v2"
)

type HTTPOrderHandler struct {
	service core.OrderService
}

func NewHTTPOrderHandler(service core.OrderService) *HTTPOrderHandler {
	return &HTTPOrderHandler{service: service}
}

func (h *HTTPOrderHandler) CreateOrder(c *fiber.Ctx) error {
	var order core.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	if err := h.service.CreateOrder(order); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})

	}
	return c.Status(fiber.StatusCreated).JSON(order)
}
