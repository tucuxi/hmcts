package handlers

import (
	"strconv"
	
	"github.com/gofiber/fiber/v2"

	"github.com/tucuxi/hmcts/internal/pkg/persistence"
)

func GetHandler(r *persistence.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		obj, exists := r.Get(id)
		if !exists {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.JSON(obj)
	}
}
