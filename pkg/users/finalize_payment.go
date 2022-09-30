package users

import (
	"bytes"
	"delta-go/pkg/common/models"
	"delta-go/pkg/common/utils"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type InputFinalizePayment struct {
	Major          string
	Faculty        string
	Year           int
	TShirtSize     string
	Diet           string
	PaymentFile    *fiber.FormFile
	FileExtension  string
	InvoiceAddress string
	FootSize       string
}

func (h handler) FinalizePayment(c *fiber.Ctx) error {

	body := new(InputFinalizePayment)
	email := c.Params("email")
	format := "2006-01-02 15:04:05"
	TimeNow := time.Now().Local()
	var placeChecker models.Place
	var user models.User
	var new_participant models.Participant

	fmt.Println(1)
	file, err := c.FormFile("PaymentFile")
	if err != nil {
		return utils.HandleResponse(c, fiber.StatusBadRequest, "Invalid PaymentFile")
	}
	fmt.Println(1)
	file_x, err := file.Open()
	if err != nil {
		return utils.HandleResponse(c, fiber.StatusBadRequest, err.Error())
	}
	fmt.Println(1)
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file_x); err != nil {
		return utils.HandleResponse(c, fiber.StatusBadRequest, err.Error())
	}

	var bytes []byte
	fmt.Println(1)
	bytes, err = json.Marshal(buf)

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
	new_participant.PaymentFile = bytes
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
