package handler

import (
	"github.com/PutskouDzmitry/golang-training-Library/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sing-up", h.signUp)
		auth.POST("/login", h.login)
		auth.GET("/refresh", h.refresh)
	}

	api := router.Group("/api/books", h.userIdentity)
	{
		api.GET("/", h.getAllBooks)
		api.GET("/:id", h.getOneBook)
		api.POST("/", h.createBook)
		api.PUT("/:id/:value", h.updateBook)
		api.DELETE("/:id", h.deleteBook)
	}
	return router
}
