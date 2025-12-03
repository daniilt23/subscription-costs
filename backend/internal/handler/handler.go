package handler

import (
	"subscription/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct{
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
	apiRouter.GET("/cost", h.GetCost)

	return r
}
