package users

import (
	"delta-go/pkg/common/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h handler) GetLeftReservations(c *fiber.Ctx) error {

	var results []map[string]interface{}
	format := "2006-01-02T15:04:05.999Z"
	TimeNow := time.Now().UTC()
	res := h.DB.Table("places").Where("(date_till_expire is null OR date_till_expire < ? ::timestamptz) AND is_sold = false", TimeNow.Format(format)).Find(&results)
	if res.Error != nil {
		return utils.HandleResponse(c, fiber.StatusInternalServerError, res.Error.Error())
	}
	return c.Status(fiber.StatusOK).JSON(len(results))
}
