package main

import (
	"delta-go/pkg/common/db"
	"delta-go/pkg/users"

	"github.com/gofiber/fiber/v2"
)

func main() {
	h := db.Init()
	app := fiber.New()

	users.RegisterRoutes(app, h)

	app.Listen("8080")
}
