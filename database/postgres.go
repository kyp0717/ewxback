package database

import (
	"log"
	"os"
  "fmt"

	"github.com/kyp0717/ewxback/model"
  "gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "gorm.io/gorm/logger"
)

var PgDBConn *gorm.DB

func PgConnectDB() {

	// Access DB credentials from environment
	host := os.Getenv("db_host")
	user := os.Getenv("db_user")
	password := os.Getenv("db_password")
	dbname := os.Getenv("db_name")
	dbport:= os.Getenv("db_port")

  fmt.Println("Starting connection with Postgres Db")
  dsn := user + "://postgres:" + password + "@" +  host + ":" + dbport + "/" + dbname + "?sslmode=disable"

  db, err := gorm.Open(postgres.Open(dsn) , &gorm.Config{})

	if err != nil {
		panic("Database connection failed.")
	}

	log.Println("Connection successful.")

	db.AutoMigrate(&model.Blog{}, &model.Item{})

	PgDBConn = db

}
