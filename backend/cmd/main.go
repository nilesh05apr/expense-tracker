package main

import (
	"github.com/gofiber/fiber/v2"
	"backend/pkg/configs"
	"backend/pkg/routes"
)


func main() {	
	app := fiber.New()

	configs.ConnectDB()

	routes.Routes(app)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Go to /api/<endpoint>... to see the API")
	})
	app.Get("/api", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello, World!"})
	})
	app.Listen(":3000")
}