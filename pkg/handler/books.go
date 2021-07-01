package handler

import (
	"fmt"
	_struct "github.com/PutskouDzmitry/golang-training-Library/pkg/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getAllBooks(c *gin.Context) {
	as, _ := c.Get("id")
	fmt.Println(as)
	books, err := h.services.Books.ReadAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, books)
}

func (h *Handler) getOneBook(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		newErrorResponse(c, http.StatusBadRequest, "error with delete book")
		return
	}
	books, err := h.services.Books.Read(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, books)
}

func (h *Handler) createBook(c *gin.Context) {
	var input _struct.Book
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Books.Add(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, id)
}
func (h *Handler) updateBook(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		newErrorResponse(c, http.StatusBadRequest, "error with delete book")
		return
	}
	value, err := strconv.Atoi(c.Param("value"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "error with delete book")
		return
	}
	_, error := h.services.Books.Update(id, value)
	if error != nil {
		newErrorResponse(c, http.StatusInternalServerError, error.Error())
	}
	c.JSON(http.StatusOK, id)
}

func (h *Handler) deleteBook(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		newErrorResponse(c, http.StatusBadRequest, "error with delete book")
		return
	}
	fmt.Println(id)
	err := h.services.Books.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, "record is deleted")
}
