package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/tucuxi/hmcts/internal/pkg/handlers"
	"github.com/tucuxi/hmcts/internal/pkg/persistence"
)

func main() {
	r := persistence.NewRepository()
	
	app := fiber.New()
	app.Get("/cases/:id", handlers.GetHandler(r))
	app.Get("/cases", handlers.ListHandler(r))
	app.Post("/cases", handlers.AddHandler(r))
	app.Delete("/cases/:id", handlers.RemoveHandler(r))
	app.Listen(":3000")
}
