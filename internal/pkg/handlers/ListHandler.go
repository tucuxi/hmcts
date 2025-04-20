package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/tucuxi/hmcts/internal/pkg/persistence"
)

func ListHandler(r *persistence.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		obj := r.List()
		return c.JSON(obj)
	}
}
