package routes

import (
	"github.com/faqih/todo_list/app/services"
	"github.com/faqih/todo_list/utils/middleware"

	"github.com/gofiber/fiber/v2"
)

// TodoRoutes contains all routes relative to /todo
func TodoRoutes(app fiber.Router) {
	r := app.Group("/todo").Use(middleware.Auth)
	// r := app.Group("/todo")

	r.Post("/create", services.CreateTodo)
	r.Get("/list", services.GetTodos)
	r.Get("/:todoID", services.GetTodo)
	r.Patch("/:todoID", services.UpdateTodoTitle)
	r.Patch("/:todoID/check", services.CheckTodo)
	r.Delete("/:todoID", services.DeleteTodo)
}
