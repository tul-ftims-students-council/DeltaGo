package users

import (
	"delta-go/pkg/common/models"
	"delta-go/pkg/common/utils"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type InputFinalizePayment struct {
	Major          string
	Faculty        string
	Year           int
	TShirtSize     string
	Diet           string
	PaymentFile    fiber.FormFile
	FileExtension  string
	InvoiceAddress string
	FootSize       string
}

func (h handler) FinalizePayment(c *fiber.Ctx) error {

	var input InputFinalizePayment
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	email := c.Params("email")
	format := "2006-01-02 15:04:05"
	TimeNow := time.Now().UTC()
	var placeChecker models.Place
	var user models.User
	var new_participant models.Participant
	var byteContainer []byte

	file, err := c.FormFile("PaymentFile")
	if err != nil {
		return utils.HandleResponse(c, fiber.StatusBadRequest, "Invalid PaymentFile")
	}

	fileContent, err := file.Open()
	defer fileContent.Close()
	if err != nil {
		return utils.HandleResponse(c, fiber.StatusBadRequest, "Invalid PaymentFile")
	}
	byteContainer, err = ioutil.ReadAll(fileContent)
	if err != nil {
		return utils.HandleResponse(c, fiber.StatusBadRequest, "Invalid PaymentFile")
	}

	if result := h.DB.Where("email = ?", email).First(&user).Error; result == gorm.ErrRecordNotFound {
		return fiber.NewError(fiber.StatusNotFound, "user not found")
	}
	if result := h.DB.Where("user_email = ?", email).First(&new_participant).Error; result != gorm.ErrRecordNotFound {
		return utils.HandleResponse(c, fiber.StatusConflict, "User already payed")
	}

	if result := h.DB.Where("user_email = ? AND is_sold = false AND date_till_expire > ?", email, TimeNow.Format(format)).First(&placeChecker).Error; result == gorm.ErrRecordNotFound {
		return utils.HandleResponse(c, fiber.StatusConflict, "User doesn't have active reservation")
	}

	fmt.Println(byteContainer)

	new_participant.UserEmail = email
	new_participant.Major = input.Major
	new_participant.Faculty = input.Faculty
	new_participant.Year = input.Year
	new_participant.TShirtSize = input.TShirtSize
	new_participant.Diet = input.Diet
	new_participant.PaymentFile = byteContainer
	new_participant.FileExtension = input.FileExtension
	new_participant.InvoiceAddress = input.InvoiceAddress
	new_participant.FootSize = input.FootSize

	if result := h.DB.Model(&placeChecker).Where("user_email = ?", email).Update("is_sold", "true").Error; result != nil {
		return utils.HandleResponse(c, fiber.StatusInternalServerError, result.Error())
	}

	if result := h.DB.Create(&new_participant); result.Error != nil {
		return utils.HandleResponse(c, fiber.StatusInternalServerError, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&user)
}
