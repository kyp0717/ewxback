package main

import (
 // "github.com/kyp0717/ewxback/dataload"
 "github.com/kyp0717/ewxback/test"
 "github.com/gofiber/fiber/v2"
 "github.com/joho/godotenv"

 "fmt"
 "log"
  // "github.com/kyp0717/ewxback/model"
  // "github.com/kyp0717/ewxback/router"
  "github.com/kyp0717/ewxback/database"
)

func init() {
	if err := godotenv.Load(".env.postgres"); err != nil {
		log.Fatal("Error in loading .env file.")
	}
	database.PgConnectDB()
  
}

func main() {
   // load data
   test.Testload(database.PgDBConn) 

   fmt.Println("Starting Fiber App")
   app := fiber.New()
   
   app.Listen(":3000")

}
