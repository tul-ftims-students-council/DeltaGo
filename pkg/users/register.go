package users

import (
	"delta-go/pkg/common/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func (h handler) Register(c *fiber.Ctx) error {
	log.Println("Registering user")
	body := models.User{}

	// parse body, attach to AddProductRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var user models.User

	user.Name = body.Name
	user.Surname = body.Surname
	user.Email = body.Email
	user.PhoneNumber = body.PhoneNumber

	if result := h.DB.Where("email = ?", body.Email).First(&user); result.Error == nil {
		return fiber.NewError(fiber.StatusConflict, "Email already exists")
	}

	// insert new db entry
	if result := h.DB.Create(&user); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&user)
}
