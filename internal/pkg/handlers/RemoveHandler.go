package handlers

import (
	"strconv"
	
	"github.com/gofiber/fiber/v2"

	"github.com/tucuxi/hmcts/internal/pkg/persistence"
)

func RemoveHandler(r *persistence.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		if !r.Remove(id) {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.SendStatus(fiber.StatusNoContent)
	}
}
