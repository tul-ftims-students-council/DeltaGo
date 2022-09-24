package main

import (
	"log"

	"delta-go/pkg/common/config"
	"delta-go/pkg/common/db"
	"delta-go/pkg/users"

	"github.com/gofiber/fiber/v2"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(&c)
	app := fiber.New()

	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	users.RegisterRoutes(app, h)

	app.Listen(c.Port)
}
