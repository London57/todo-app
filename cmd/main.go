package main

import (
	"fmt"
	"log"

	"github.com/London57/todo-app/config"
	"github.com/London57/todo-app/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("config error: %w", err))
	}

	app.Run(cfg)
}
