package handler

import (
	"github.com/PutskouDzmitry/golang-training-Library/pkg/entity"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input entity.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type singInInput struct {
	Id       primitive.ObjectID `bson:"_id"`
	Username string             `json:"username" binding:"required"`
	Password string             `json:"password" binding:"required"`
}

func (h *Handler) login(c *gin.Context) {
	var input singInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	accessToken, err := h.services.Authorization.GenerateTokenAccessToken(input.Id, input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	refreshToken, err := h.services.Authorization.GenerateTokenRefreshToken(input.Id, input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.SetCookie("refreshToken", refreshToken, 20, "/", "localhost", true, true)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": accessToken,
	})
}

func (h *Handler) refresh(c *gin.Context) {
	var input singInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	accessToken, err := h.services.Authorization.GenerateTokenAccessToken(input.Id, input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	refreshToken, err := h.services.Authorization.GenerateTokenRefreshToken(input.Id, input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.SetCookie("refreshToken", refreshToken, 20, "/", "localhost", true, true)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": accessToken,
	})
}
