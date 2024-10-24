package controller

import (
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/kyp0717/ewxback/database"
	"github.com/kyp0717/ewxback/webapp/model"
)

// Item list
func ItemList(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Item List",
	}

	// Sleep to add some delay in API response
	time.Sleep(time.Millisecond * 1500)

	db := database.PgDBConn

	var records []model.Item

	db.Find(&records)

	context["item_records"] = records

	c.Status(200)
	return c.JSON(context)
}

