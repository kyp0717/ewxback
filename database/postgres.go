package database

import (
	"log"
	"os"
  "fmt"

  "database/sql"
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
  // connStr := "postgres://postgres:password@localhost:5432/tbl_ews?sslmode=disable"
  dsn := user + "://postgres:" + password + "@" +  host + ":" + dbport + "/" + dbname + "?sslmode=disable"

  pgDB, err := sql.Open("postgres", dsn)

  db, err := gorm.Open(postgres.New(postgres.Config{
    Conn: pgDB,
  }), &gorm.Config{})


	if err != nil {
		panic("Database connection failed.")
	}

	log.Println("Connection successful.")

	db.AutoMigrate(new(model.Blog))

	PgDBConn = db
}
