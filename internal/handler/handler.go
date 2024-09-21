package handler

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/GoSeoTaxi/dh9_dProxy/internal/config"
	"github.com/GoSeoTaxi/dh9_dProxy/internal/service"
)

type Handler struct {
	service *service.Service
	logger  *zap.Logger
	config  *config.Config
}

func NewHandler(s *service.Service, logger *zap.Logger, cfg *config.Config) *Handler {
	return &Handler{service: s, logger: logger, config: cfg}
}

func (h *Handler) GetNew(w http.ResponseWriter, r *http.Request) {
	_ = r
	h.logger.Info("Получен запрос на /get_new")

	result, err := h.service.HandleGetNew()
	if err != nil {
		h.logger.Error("Ошибка при выдаче прокси /get_new", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.logger.Info("Запрос на /get_new успешно обработан")
	_, _ = w.Write([]byte(result + "\n"))
}

func Router(lc fx.Lifecycle, handler *Handler) http.Handler {
	r := chi.NewRouter()

	r.Get("/get_new", handler.GetNew)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			addr := ":" + handler.config.ServerPort
			handler.logger.Info("Запуск HTTP сервера", zap.String("address", addr))
			go func() { _ = http.ListenAndServe(addr, r) }()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			handler.logger.Info("Остановка HTTP сервера")
			return nil
		},
	})

	return r
}
