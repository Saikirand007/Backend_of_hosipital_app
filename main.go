package main

import (
	"fmt"

	"github.com/Saikirand007/Backend_of_hosipital_app/database"
	"github.com/Saikirand007/Backend_of_hosipital_app/hospital"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

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

func initDataBase() {

	var err error
	database.DBConn, err = gorm.Open("sqlite3", "hosipitalRecod.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")

}

func main() {
	app := fiber.New()
	initDataBase()

	setupRoutes(app)

	app.Listen(":3000")

	defer database.DBConn.Close()
}
