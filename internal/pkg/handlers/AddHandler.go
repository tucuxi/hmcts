package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/tucuxi/hmcts/internal/pkg/persistence"
	"github.com/tucuxi/hmcts/pkg/domain"
)

func AddHandler(r *persistence.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var obj domain.Case
		if err := c.BodyParser(&obj); err != nil {
			return err
		}
		r.Add(obj)
		return c.JSON(obj)
	}
}
