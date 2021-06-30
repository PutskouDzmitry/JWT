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
		auth.POST("/login", h.signIn)
		auth.GET("/refresh", h.refresh)
	}

	api := router.Group("/api", h.userIdentity)
	{
		api.GET("/books", h.getAllBooks)
		api.GET("/book{id}", h.getOneBook)
		api.POST("/books", h.createBook)
		api.PUT("/books{id}/{number}", h.updateBook)
		api.DELETE("/books{id}", h.deleteBook)
	}
	return router
}