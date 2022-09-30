package users

import (
	"delta-go/pkg/common/models"
	"delta-go/pkg/common/utils"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type InputFinalizePayment struct {
	Major          string `validate:"required,min=4"`
	Faculty        string `validate:"required"`
	Year           int    `validate:"required"`
	TShirtSize     string `validate:"required"`
	Diet           string `validate:"required"`
	PaymentFile    []byte `validate:"required"`
	FileExtension  string `validate:"required,min=3"`
	InvoiceAddress string
	FootSize       string `validate:"required"`
}

func (h handler) FinalizePayment(c *fiber.Ctx) error {

	body := new(InputFinalizePayment)
	email := c.Params("email")
	format := "2006-01-02 15:04:05"
	TimeNow := time.Now().Local()
	var placeChecker models.Place
	var user models.User
	var new_participant models.Participant

	if err := c.BodyParser(&body); err != nil {
		return utils.HandleResponse(c, fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return utils.HandleResponse(c, fiber.StatusBadRequest, err.Error())
	}

	if result := h.DB.Where("email = ?", email).First(&user).Error; result == gorm.ErrRecordNotFound {
		return fiber.NewError(fiber.StatusNotFound, "user not found")
	}

	if result := h.DB.Where("user_email = ? AND is_sold = false AND date_till_expire > ?", email, TimeNow.Format(format)).First(&placeChecker).Error; result == gorm.ErrRecordNotFound {
		return utils.HandleResponse(c, fiber.StatusConflict, "User doesn't have active reservation")
	}

	new_participant.UserEmail = email
	new_participant.Major = body.Major
	new_participant.Faculty = body.Faculty
	new_participant.Year = body.Year
	new_participant.TShirtSize = body.TShirtSize
	new_participant.Diet = body.Diet
	new_participant.PaymentFile = body.PaymentFile
	new_participant.FileExtension = body.FileExtension
	new_participant.InvoiceAddress = body.InvoiceAddress
	new_participant.FootSize = body.FootSize

	if result := h.DB.Model(&placeChecker).Where("user_email = ?", email).Update("is_sold", "true").Error; result != nil {
		return utils.HandleResponse(c, fiber.StatusInternalServerError, result.Error())
	}

	if result := h.DB.Create(&new_participant); result.Error != nil {
		return utils.HandleResponse(c, fiber.StatusInternalServerError, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&user)
}
