package users

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := app.Group("/users")
	routes.Post("/register", h.Register)
	routes.Get("/:email", h.GetUser)
	routes.Get("/:email/payment/start", h.StartPayment)
	routes.Post("/:email/payment/send", h.FinalizePayment)
}
