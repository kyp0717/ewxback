package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
  "github.com/kyp0717/ewxback/router"
  "github.com/kyp0717/ewxback/database"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error in loading .env file.")
	}

	database.ConnectDB()
}

func main() {

	sqlDb, err := database.DBConn.DB()

	if err != nil {
		panic("Error in sql connection.")
	}

	defer sqlDb.Close()

	app := fiber.New()

	app.Static("/static", "./static")

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(logger.New())

	router.SetupRoutes(app)

	app.Listen(":8000")
}
