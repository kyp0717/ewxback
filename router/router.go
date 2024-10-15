package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kyp0717/ewxback/controller"
)

// Setup routing information
func SetupRoutes(app *fiber.App) {

	// list => get
	// read blog => get (id)
	// add => post
	// update => put
	// delete => delete

	app.Get("/blog", controller.BlogList)
	app.Get("/:id", controller.BlogDetail)
	app.Post("/", controller.BlogCreate)
	app.Put("/:id", controller.BlogUpdate)
	app.Delete("/:id", controller.BlogDelete)

	app.Get("/test", controller.TestList)
}
