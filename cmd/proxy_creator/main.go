package main

import (
	"github.com/GoSeoTaxi/dh9_dProxy/internal/config"
	"github.com/GoSeoTaxi/dh9_dProxy/internal/handler"
	"github.com/GoSeoTaxi/dh9_dProxy/internal/repository"
	"github.com/GoSeoTaxi/dh9_dProxy/internal/service"

	"net/http"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	app := fx.New(
		fx.Provide(
			config.NewConfig,
			repository.NewRepository,
			service.NewService,
			handler.NewHandler,
			handler.Router,
			NewLogger,
		),
		fx.Invoke(func(r http.Handler) {}),
	)

	app.Run()
}

func NewLogger() (*zap.Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	defer func() { _ = logger.Sync() }()
	return logger, nil
}
