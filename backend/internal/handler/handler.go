package handler

import (
	"subscription/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "subscription/docs"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()

	apiRouter := r.Group("/api")

	apiRouter.POST("/create", h.CreateSubscription)
	apiRouter.POST("/cost", h.GetCost)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
