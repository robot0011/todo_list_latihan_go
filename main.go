package main

import (
	"log/slog"
	"os"

	"github.com/faqih/todo_list/config/database"
)

func main() {
	// Run your server.
	database.Connect()
	// database.Migrate(&dal.User{}, &dal.Todo{})

	if err := runServer(); err != nil {
		slog.Error("Failed to start server!", "details", err.Error())
		os.Exit(1)
	}
}
