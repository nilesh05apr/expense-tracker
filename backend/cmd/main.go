package main

import (
	"github.com/gofiber/fiber/v2"
cf	"backend/pkg/configs" 
)


func main() {	
	app := fiber.New()

	cf.ConnectDB()
	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/api", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello, World!"})
	})
	app.Listen(":3000")
}