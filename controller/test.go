package controller

import (
	"time"
  "fmt"
  "encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/kyp0717/ewxback/database"
	"github.com/kyp0717/ewxback/model"
)

// Item list
func TestList2() error {
	// Sleep to add some delay in API response
	time.Sleep(time.Millisecond * 3500)

	db := database.PgDBConn
	var ttt []model.TestTable

  db.Find(&ttt)
  t,e := json.Marshal(ttt)
  fmt.Println(string(t))
  return e
}


// Item list
func TestList(c *fiber.Ctx) error {
	// Sleep to add some delay in API response
	time.Sleep(time.Millisecond * 3500)

	db := database.PgDBConn
	var ttt []model.TestTable

  db.Find(&ttt)

  if len(records) == 0 {
    return c.Status(404).JSON(fiber.Map{"status":"error", "message":"No record", "data": nil})
  }

	context["test_records"] = records
  return c.Status(200).JSON(fiber.Map{"status":"success", "message":"record found", "data": records})
 
}

func BlogListTest(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Blog List",
	}

	// Sleep to add some delay in API response
	time.Sleep(time.Millisecond * 1500)

	db := database.PgDBConn

	var records []model.Blog

	db.Find(&records)

	context["blog_records"] = records

	c.Status(200)
	return c.JSON(context)
}
