package handler

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	userId, err := h.services.ParseAccessToken(headerParts[1])
	if err != nil {
		refreshToken, err := c.Cookie("refreshToken")
		if err != nil {
			newErrorResponse(c, http.StatusUnauthorized, err.Error())
		}
		_, err = h.services.ParseRefreshToken(refreshToken)
		if err != nil {
			newErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}
		id, _ := primitive.ObjectIDFromHex(userId)
		accessToken, err := h.services.GenerateTokenAccessToken(id, "test", "test")
		if err != nil {
			newErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}
		c.JSON(http.StatusOK, map[string]interface{}{
			"Warn": "Please, choose login for update your tokens",
		})
		c.JSON(http.StatusOK, map[string]interface{}{
			"id": accessToken,
		})
	}
	c.Set("userId", userId)
}
