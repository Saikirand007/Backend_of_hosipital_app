package hospital

import (
	"github.com/gofiber/fiber"
)

func GetPatientsDetails(c *fiber.Ctx) {

	c.Send("ALL Patients are ravi, sreekanth, hemantha")

}

func GetBookingIds(c *fiber.Ctx) {
	c.Send("ALL Booking Ids are 11453, 1225,133")

}

func NewBookingIds(c *fiber.Ctx) {

	c.Send("All New Patients Ids are none")

}
func DeleteOldBookingIds(c *fiber.Ctx) {

	c.Send("Delete Old booking Ids will be None")

}
