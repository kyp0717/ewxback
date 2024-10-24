package main

import (
 //"github.com/kyp0717/ewxback/test"

   "github.com/gofiber/fiber/v2"
   "github.com/joho/godotenv"
   "github.com/gofiber/fiber/v2/middleware/logger"
   "fmt"
   "log"
   "github.com/kyp0717/ewxback/webapp/router"
   "github.com/kyp0717/ewxback/database"

 // "github.com/kyp0717/ewxback/webapp/controller"
)

func init() {
	if err := godotenv.Load(".env.postgres"); err != nil {
		log.Fatal("Error in loading .env file.")
	}
	database.PgConnectDB()
  
}

func main() {




   fmt.Println("Starting Fiber App")
   app := fiber.New()
   app.Use(logger.New())
	router.SetupRoutes(app) 
   app.Listen(":3000")

}
