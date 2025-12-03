package app

import (
	"log"
	"subscription/internal/config"
	"subscription/internal/handler"
	"subscription/internal/repository/postgres"
	"subscription/internal/repository/postgres/subscription"
	"subscription/internal/service"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) Init() {
	log.Println("START INIT APP")
	cfg := config.MustLoad()

	db := postgres.Connect(cfg)
	defer db.Close()
	log.Println("Connect to database")
	subRepo := subscription.NewSubscriptionRepoSQL(db)

	service := service.NewService(subRepo)

	handler := handler.NewHandler(service)
	routes := handler.InitRoutes()

	server := NewServer(cfg, routes)

	if err := server.Run(); err != nil {
		log.Fatalf("cannot start server %v", err.Error())
	}
}
