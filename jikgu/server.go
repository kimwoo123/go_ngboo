package main

import (
	"jikgu/router"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()
	router.SetupRoutes(app)
	log.Fatal(app.Listen(":4223"))
}
