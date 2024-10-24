package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kyp0717/ewxback/webapp/controller"
)

// Setup routing information
func SetupRoutes(app *fiber.App) {
    
	// list => get
	// read blog => get (id)
	// add => post
	// update => put
	// delete => delete
	app.Get("/", handler.index)
	app.Get("/blog", controller.BlogList)
    app.Get("/test", controller.TestList)
}
