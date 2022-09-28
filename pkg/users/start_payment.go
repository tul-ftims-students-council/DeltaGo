package users

import (
	"delta-go/pkg/common/models"
	"delta-go/pkg/common/utils"
	"fmt"
	"reflect"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Search in places for a first place that has no reservation made yet.
// but if there are no free places look if any of reservation is expired.
// if there is any overwrite it with new timestamp and user email address
// but if all of the reservations are not expired return no places are left.
func (h handler) StartPayment(c *fiber.Ctx) error {
	email := c.Params("email")
	fmt.Println("Starting payment on a user ", email)

	if email == "" || reflect.TypeOf(email).Kind() != reflect.String {
		return fiber.NewError(fiber.StatusBadRequest, "email must be a string")
	}
	// check if user exists
	var user models.User

	if result := h.DB.Where("email = ?", email).First(&user).Error; result == gorm.ErrRecordNotFound {
		return fiber.NewError(fiber.StatusNotFound, "user not found")
	}
	var placeChecker models.Place
	format := "2006-01-02 15:04:05"
	TimeNow := time.Now().Local()
	Time := time.Now().Local().Add(time.Hour * 4).Add(time.Minute * 20)

	if result := h.DB.Where("user_email = ? AND is_sold = false AND date_till_expire > ?", email, TimeNow.Format(format)).First(&placeChecker); result != nil && result.Error != gorm.ErrRecordNotFound {
		fmt.Println(result)
		return utils.HandleResponse(c, fiber.StatusConflict, "User already started payment")
	}

	if result := h.DB.Where(map[string]interface{}{"user_email": email, "is_sold": true}).First(&placeChecker); result != nil && result.Error != gorm.ErrRecordNotFound {
		return utils.HandleResponse(c, fiber.StatusConflict, "User already bought a place")
	}

	fmt.Println(Time.Format(format))
	//Find if there is any free place
	var freePlace models.Place
	if result := h.DB.Where("(user_email is null OR date_till_expire is null OR date_till_expire <= ? ::date) AND is_sold = false", Time.Format(format)).First(&freePlace); result.Error != nil {
		return utils.HandleResponse(c, fiber.StatusOK, "There are no free places left")
	}
	// Update free place
	result := h.DB.Model(models.Place{}).Where("id = ?", freePlace.ID).Updates(models.Place{UserEmail: user.Email, DateTillExpire: Time})
	if result.Error != nil {
		return utils.HandleResponse(c, fiber.StatusInternalServerError, result.Error.Error())
	}
	return utils.HandleResponse(c, fiber.StatusOK, string(Time.Format(format)))
}
