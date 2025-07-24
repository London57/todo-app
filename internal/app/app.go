package app

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/London57/todo-app/config"
	"github.com/London57/todo-app/internal/controller/http"
	"github.com/London57/todo-app/internal/controller/http/common"
	v1 "github.com/London57/todo-app/internal/controller/http/v1"
	"github.com/London57/todo-app/internal/controller/http/v1/auth"
	"github.com/London57/todo-app/internal/controller/http/v1/item"
	"github.com/London57/todo-app/internal/controller/http/v1/list"
	"github.com/London57/todo-app/internal/repo/persistent"
	"github.com/London57/todo-app/internal/usecase/signup"
	"github.com/London57/todo-app/pkg/httpserver"
	"github.com/London57/todo-app/pkg/logger"
	"github.com/London57/todo-app/pkg/postgres"
	"github.com/go-playground/validator/v10"
)

func Run(cfg *config.Config) {
	var l *logger.Log
	switch cfg.App.Mode {
	case "prod":
		l = logger.New(
			os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo},
		)
	default:
		l = logger.New(
			os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug},
		)
	}
	// postgresql://<user>:<password>@<host>:<port>/<database>?<params>
	db := cfg.DB
	pg, err := postgres.New(
		l,
		fmt.Sprintf(
			"postgres://%s:%s@%s:%d/%s?sslmode=%s",
			db.User,
			db.Password,
			db.Host,
			db.Port,
			db.DataBase,
			db.SSLMode,
		),
	)
	if err != nil {
		l.Fatal(fmt.Errorf("Run - postgres.New: %w", err).Error())
	}
	defer pg.Close()

	httpserver := httpserver.New(httpserver.Address(cfg.API.Host, cfg.API.Port))

	userRepo := persistent.New(pg)

	signupUC := signup.New(userRepo)

	validator := validator.New()

	bC := common.New(l, validator)
	authC := auth.NewAuthController(bC, &signupUC, cfg)
	listC := list.New(bC)
	itemC := item.New(bC)
	v1 := v1.New(authC, listC, itemC, cfg)

	http.NewRouter(httpserver.App, &v1, cfg)
	httpserver.HTTPServer.Handler = httpserver.App

	httpserver.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-interrupt:
		l.Warn(fmt.Sprintf("app - Run - signals: %s", err.String()))

	case err := <-httpserver.Notify():
		l.Error(fmt.Errorf("app - Run - httpserver: %w", err).Error())
	}

	err = httpserver.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err).Error())
	}

}
