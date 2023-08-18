package main

import (
	"os"

	"github.com/akramfirmansyah/go-react/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func init() {
	utils.LoadEnv()
}

func main() {
	// Load templates
	engine := html.New("./views", ".tmpl")

	// Create app
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Config App
	app.Static("/", "./public")

	// Routing
	clientRoutes := []string{
		"/",
		"/about",
	}

	for _, route := range clientRoutes {
		app.Get(route, func(c *fiber.Ctx) error {
			return c.Render("index", nil)
		})
	}

	app.Listen(":" + os.Getenv("PORT"))
}
