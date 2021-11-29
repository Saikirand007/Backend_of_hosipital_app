package main

import (
	"/home/debasish/Desktop/Hosipital_Backend_api/hospital/hospital"

	"github.com/gofiber/fiber"
)

func helloWorld(c *fiber.Ctx) {

	c.Send("Hello,World")
}

func setupRoutes(app *fiber.App) {

	app.Get("/", helloWorld)

	app.Get("/api/v1/hospital", hospital.getpatientsDetails)
	app.Get("/api/v1/hospital", hospital.getbookingIds)
	app.Get("/api/v1/hospital", hospital.NewBookingIds)
	app.Get("/api/v1/hospital", hospital.deleteoldbookingIds)
}

func main() {
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
