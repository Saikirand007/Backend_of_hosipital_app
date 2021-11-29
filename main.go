package main

import (
	"github.com/Saikirand007/Backend_of_hosipital_app/hospital"

	"github.com/gofiber/fiber"
)

func helloWorld(c *fiber.Ctx) {

	c.Send("Hello,World")
}

func setupRoutes(app *fiber.App) {

	app.Get("/", helloWorld)

	app.Get("/api/v1/hospital", hospital.GetPatientsDetails)
	app.Get("/api/v1/hospital", hospital.GetBookingIds)
	app.Get("/api/v1/hospital", hospital.NewBookingIds)
	app.Get("/api/v1/hospital", hospital.DeleteOldBookingIds)
}

func main() {
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
