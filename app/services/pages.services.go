package services

import (
	// "errors"

	// "github.com/faqih/todo_list/app/dal"
	// "github.com/faqih/todo_list/app/types"
	// "github.com/faqih/todo_list/utils"

	// "github.com/gofiber/fiber/v2"
	"log"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/session"
	// "github.com/gofiber/template/html"
	// "gorm.io/gorm"
)

func RenderLoginPage(c *fiber.Ctx) error {
	if err := c.Render("login", fiber.Map{}); err != nil {
		log.Println("Error rendering template:", err)
		return err
	}
	return nil
}
func RenderMainPage(c *fiber.Ctx) error {
	if err := c.Render("main", fiber.Map{}); err != nil {
		log.Println("Error rendering template:", err)
		return err
	}
	return nil
}

func RenderSignPage(c *fiber.Ctx) error {
	return c.Render("signup", fiber.Map{})
}
func RenderTodoPage(c *fiber.Ctx) error {
	return c.Render("todos", fiber.Map{})
}
