package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/faqih/todo_list/app/routes"

	// "github.com/faqih/todo_list/config"
	"github.com/faqih/todo_list/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"

	gowebly "github.com/gowebly/helpers"
)

// runServer runs a new HTTP server with the loaded environment variables.
func runServer() error {
	// Validate environment variables.
	port, err := strconv.Atoi(gowebly.Getenv("BACKEND_PORT", "8000"))
	if err != nil {
		return err
	}

	// Create a new server instance with options from environment variables.
	// For more information, see https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
	config := fiber.Config{
		Views: html.NewFileSystem(http.Dir("./templates"), ".html"),
		// ViewsLayout:  "main",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorHandler: utils.ErrorHandler,
	}

	// Create a new Fiber server.
	server := fiber.New(config)

	routes.AuthRoutes(server)
	routes.TodoRoutes(server)
	routes.PageRoutes(server)
	server.Static("/static", "./static")
	server.Static("/templates", "./templates")

	// server.Get("/", func(c *fiber.Ctx) error {
	// 	return c.Render("login", nil)
	// })

	// Add Fiber middlewares.
	server.Use(logger.New())

	// Handle static files.
	// server.Static("/static", "./static")

	// Handle index page view.
	// server.Get("/", indexViewHandler)

	// // Handle API endpoints.
	// server.Get("/api/hello-world", showContentAPIHandler)

	return server.Listen(fmt.Sprintf(":%d", port))
}
