package main

import (
 "github.com/kyp0717/ewxback/test"
 "github.com/gofiber/fiber/v2"
 "github.com/joho/godotenv"
 "github.com/gofiber/fiber/v2/middleware/logger"
 "fmt"
 "log"
  "github.com/kyp0717/ewxback/router"
  "github.com/kyp0717/ewxback/database"
  "github.com/kyp0717/ewxback/controller"
)

func init() {
	if err := godotenv.Load(".env.postgres"); err != nil {
		log.Fatal("Error in loading .env file.")
	}
	database.PgConnectDB()
  
}

func main() {
   // load data
   // test.Testload1(database.PgDBConn) 
   // test.Testload2(database.PgDBConn) 

   controller.PrintList() 

   fmt.Println("Starting Fiber App")
   app := fiber.New()
   app.Use(logger.New())
	 router.SetupRoutes(app) 
   app.Listen(":3000")

}
