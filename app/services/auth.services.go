package services

import (
	"errors"
	"fmt"

	"github.com/faqih/todo_list/app/dal"
	"github.com/faqih/todo_list/app/types"
	"github.com/faqih/todo_list/utils"
	"github.com/faqih/todo_list/utils/jwt"
	"github.com/faqih/todo_list/utils/password"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// LogContext logs details of the Fiber context for debugging
// LogContext logs details of the Fiber context for debugging
func LogContext(ctx *fiber.Ctx) {
	fmt.Println("Request Information:")
	fmt.Printf("Method: %s\n", ctx.Method())
	fmt.Printf("Path: %s\n", ctx.Path())
	fmt.Printf("Host: %s\n", ctx.Hostname())
	fmt.Printf("Protocol: %s\n", ctx.Protocol())
	fmt.Printf("IP: %s\n", ctx.IP())
	fmt.Printf("IPs: %v\n", ctx.IPs())
	fmt.Printf("Original URL: %s\n", ctx.OriginalURL())
	fmt.Printf("Headers: %v\n", ctx.GetReqHeaders())

	// Print the request body (only if small or if logging is secure)
	fmt.Printf("Body: %s\n", string(ctx.Body()))
}

// Login service logs in a user
func Login(ctx *fiber.Ctx) error {
	// Log the context details
	// LogContext(ctx)

	b := new(types.LoginDTO)

	if err := utils.ParseBodyAndValidate(ctx, b); err != nil {
		return err
	}

	u := &types.UserResponse{}

	err := dal.FindUserByEmail(u, b.Email).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}

	if err := password.Verify(u.Password, b.Password); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}

	t := jwt.Generate(&jwt.TokenPayload{
		ID: u.ID,
	})

	return ctx.JSON(&types.AuthResponse{
		User: u,
		Auth: &types.AccessResponse{
			Token: t,
		},
	})
}

// Signup service creates a user
func Signup(ctx *fiber.Ctx) error {
	// Log the context details
	LogContext(ctx)

	b := new(types.SignupDTO)

	if err := utils.ParseBodyAndValidate(ctx, b); err != nil {
		return err
	}

	err := dal.FindUserByEmail(&struct{ ID string }{}, b.Email).Error

	// If email already exists, return
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusConflict, "Email already exists")
	}

	user := &dal.User{
		Name:     b.Name,
		Password: password.Generate(b.Password),
		Email:    b.Email,
	}

	// Create a user, if error return
	if err := dal.CreateUser(user); err.Error != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error.Error())
	}

	// generate access token
	t := jwt.Generate(&jwt.TokenPayload{
		ID: user.ID,
	})

	return ctx.JSON(&types.AuthResponse{
		User: &types.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
		Auth: &types.AccessResponse{
			Token: t,
		},
	})
}
