package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Static("/public", "./public")
	app.Get("/", mainPage)
	app.Listen(":3000")
}

func mainPage(c *fiber.Ctx) error {
	//This function will be see different soon
	return c.SendString("Hellö")
}
