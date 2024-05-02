package router

import (
	"jikgu/handler"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)
	//Address

	//Name
	name := api.Group("/romanization")
	name.Get("/", handler.GetName)
}
