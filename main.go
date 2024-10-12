package main

import (
 "github.com/gofiber/fiber/v2"
 "github.com/joho/godotenv"

 // "database/sql"
 "fmt"
 // _ "github.com/lib/pq"
 "log"
  // "gorm.io/driver/postgres"
  // "gorm.io/gorm"
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

   fmt.Println("Starting Fiber App")
   app := fiber.New()

   // connStr := "postgres://postgres:password@localhost:5432/tbl_ews?sslmode=disable"
   // pgDB, err := sql.Open("postgres", connStr)
   //
   // db, err := gorm.Open(postgres.New(postgres.Config{
   //   Conn: pgDB,
   // }), &gorm.Config{})
   //
   // defer pgDB.Close()
   // if err != nil {
   //  log.Fatal(err)
   // }
   // if err = pgDB.Ping(); err != nil {
   //  log.Println("DB Ping Failed")
   //  log.Fatal(err)
   // }
   // log.Println("DB Connection started successfully")
   //
   // db.Migrator().CreateTable(&model.Blog{})

   
   app.Listen(":3000")

}
