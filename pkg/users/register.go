package users

import (
	"delta-go/pkg/common/models"
	"delta-go/pkg/common/utils"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Input struct {
	Name        string `validate:"required,min=4"`
	Surname     string `validate:"required,min=4"`
	Email       string `validate:"required,min=4"`
	PhoneNumber string `validate:"required,min=4"`
}

func (h handler) Register(c *fiber.Ctx) error {
	fmt.Println("Registering user")
	body := new(Input)

	if err := c.BodyParser(&body); err != nil {
		return utils.HandleResponse(c, fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return utils.HandleResponse(c, fiber.StatusBadRequest, err.Error())
	}

	var user models.User

	user.Name = body.Name
	user.Surname = body.Surname
	user.Email = body.Email
	user.PhoneNumber = body.PhoneNumber

	if result := h.DB.Where("Email = ?", body.Email).First(&user); result.Error == nil {
		return utils.HandleResponse(c, fiber.StatusConflict, "Użytkownik z takim mailem już istnieje")
	}

	if result := h.DB.Create(&user); result.Error != nil {
		return utils.HandleResponse(c, fiber.StatusInternalServerError, result.Error.Error())
	}

	if err := utils.SendEmail(body.Email, "Rejestracja", "Rejestracja przebiegła pomyślnie"); err != nil {
		fmt.Println(err)
	}

	return utils.HandleResponse(c, fiber.StatusOK, "Użytkownik został zarejestrowany")
}
