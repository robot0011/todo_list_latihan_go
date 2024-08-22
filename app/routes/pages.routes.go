package routes

import (
	"github.com/faqih/todo_list/app/services"

	"github.com/gofiber/fiber/v2"
)

// TodoRoutes contains all routes relative to /todo
func PageRoutes(app fiber.Router) {
	// r := app.Group("/todo")
	app.Get("/", services.RenderLoginPage)
	// app.Get("/", services.RenderMainPage)
	app.Get("/signup", services.RenderSignPage)
	app.Get("/content", services.RenderTodoPage)

	// r.Post("/create", services.CreateTodo)
	// r.Get("/list", services.GetTodos)
	// r.Get("/:todoID", services.GetTodo)
	// r.Patch("/:todoID", services.UpdateTodoTitle)
	// r.Patch("/:todoID/check", services.CheckTodo)
	// r.Delete("/:todoID", services.DeleteTodo)
}
