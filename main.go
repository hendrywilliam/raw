package main

import "github.com/gofiber/fiber/v3"

func main() {
	app := fiber.New()
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("hi")
	})
	app.Listen(":8080")
}