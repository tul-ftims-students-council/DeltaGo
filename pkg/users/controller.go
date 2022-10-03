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

	routes_user := app.Group("/users")
	routes_reservations := app.Group("/reservations")
	routes_reservations.Get("/left", h.GetLeftReservations)
	routes_user.Post("/register", h.Register)
	routes_user.Get("/:email", h.GetUser)
	routes_user.Get("/:email/payment/start", h.StartPayment)
	routes_user.Post("/:email/payment/send", h.FinalizePayment)
}
