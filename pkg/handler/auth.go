package handler

import (
	"github.com/PutskouDzmitry/golang-training-Library/pkg/struct"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)


func (h *Handler) signUp(c *gin.Context) {
	var input _struct.User
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
		"id":id,
	})
}

type singInInput struct {
	Id primitive.ObjectID `bson:"_id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
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
	_, err = h.services.Authorization.GenerateTokenRefreshToken(input.Id, input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id":accessToken,
	})
}

func (h *Handler) refresh(c *gin.Context) {

}

