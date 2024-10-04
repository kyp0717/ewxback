package main

import (
 "github.com/gofiber/fiber/v2"
 "database/sql"
 "fmt"
 _ "github.com/lib/pq"
 "log"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)
type User struct {
    Name string
}

func main() {

   app := fiber.New()

   fmt.Println("Starting connection with Postgres Db")
   connStr := "postgres://postgres:password@localhost:5432/tbl_ews?sslmode=disable"
   pgDB, err := sql.Open("postgres", connStr)

   db, err := gorm.Open(postgres.New(postgres.Config{
     Conn: pgDB,
   }), &gorm.Config{})

   defer pgDB.Close()
   if err != nil {
    log.Fatal(err)
   }
   if err = pgDB.Ping(); err != nil {
    log.Println("DB Ping Failed")
    log.Fatal(err)
   }
   log.Println("DB Connection started successfully")

   // Create table for `User`
   db.Migrator().CreateTable(&User{})

   
   app.Listen(":3000")

}
