package app

import (
	"log"
	"subscription/internal/config"
	"subscription/internal/handler"
	"subscription/internal/logger"
	"subscription/internal/repository/postgres"
	"subscription/internal/repository/postgres/subscription"
	"subscription/internal/service"

	"go.uber.org/zap"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) Init() {
	log.Println("START INIT APP")
	cfg := config.MustLoad()

	logger := logger.InitLogger()
	defer func() {
		err := logger.Sync()
		if err != nil {
			logger.Error("failed to sync logger", zap.Error(err))
		}
	}()
	logger.Info("Init zap logger")

	db := postgres.Connect(cfg, logger)
	defer func() {
		err := db.Close()
		if err != nil {
			logger.Error("Cannot clode db connection")
		}
	}()

	subRepo := subscription.NewSubscriptionRepoSQL(db)
	logger.Info("Init subscription repository")

	service := service.NewService(subRepo, logger)
	logger.Info("Init subscription service")

	handler := handler.NewHandler(service)
	logger.Info("Init subscription handlers")
	routes := handler.InitRoutes()
	logger.Info("Init subscription routes")

	server := NewServer(cfg, routes)

	err := server.Run()
	if err != nil {
		logger.Error("Error", zap.String("Cannot run server", err.Error()))
	}
}
