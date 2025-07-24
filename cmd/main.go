package main

import (
	"github.com/London57/todo-app/config"
	"github.com/London57/todo-app/internal/app"
)

func main() {
	cfg := config.NewConfig()

	app.Run(cfg)
}
