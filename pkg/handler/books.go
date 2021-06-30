package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getAllBooks(c *gin.Context) {
	id, ok := c.Get("userId")
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}
	fmt.Println(id)
	//users, err := a.Data.ReadAll()
	//if err != nil {
	//	_, err := writer.Write([]byte("got an error when tried to get users "))
	//	if err != nil {
	//		log.Println(err)
	//	}
	//}
	//logrus.Info(users)
	//err = json.NewEncoder(writer).Encode(users)
	//if err != nil {
	//	log.Println(err)
	//	writer.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
}

func (h *Handler) getOneBook(c *gin.Context) {
	//idRequest := mux.Vars(request)
	//id := idRequest["id"]
	//user, err := a.Data.Read(id)
	//if err != nil {
	//	_, err := writer.Write([]byte("got an error when tried to get users"))
	//	if err != nil {
	//		log.Println(err)
	//		writer.WriteHeader(http.StatusInternalServerError)
	//		return
	//	}
	//}
	//logrus.Info(user)
	//if user.NameOfBook != "" {
	//	err = json.NewEncoder(writer).Encode(user)
	//	if err != nil {
	//		log.Println(err)
	//		writer.WriteHeader(http.StatusInternalServerError)
	//		return
	//	}
	//}
}

func (h *Handler) createBook(c *gin.Context) {
	//book := new(data.Book)
	//err := json.NewDecoder(request.Body).Decode(&book)
	//if err != nil {
	//	log.Printf("failed reading JSON: %s", err)
	//	writer.WriteHeader(http.StatusBadRequest)
	//	return
	//}
	//if book == nil {
	//	log.Printf("failed empty JSON")
	//	writer.WriteHeader(http.StatusBadRequest)
	//	return
	//}
	//err = a.Data.Add(*book)
	//err = json.NewEncoder(writer).Encode(book.BookId)
	//if err != nil {
	//	log.Println(err)
	//	writer.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//if err != nil {
	//	log.Println("user hasn't been created")
	//	writer.WriteHeader(http.StatusBadRequest)
	//	return
	//}
	//writer.WriteHeader(http.StatusCreated)
}
func (h *Handler) updateBook(c *gin.Context) {
	//idRequest := mux.Vars(request)
	//id := idRequest["id"]
	//strNumber := idRequest["number"]
	//number, err := strconv.Atoi(strNumber)
	//if err != nil {
	//	log.Println("book hasn't been updated, because number doesn't equal int:", number)
	//	writer.WriteHeader(http.StatusBadRequest)
	//	return
	//}
	//err = a.Data.Update(id, number)
	//if err != nil {
	//	log.Println("book hasn't been updated")
	//	writer.WriteHeader(http.StatusBadRequest)
	//	return
	//}
	//err = json.NewEncoder(writer).Encode(id)
	//if err != nil {
	//	log.Println(err)
	//	writer.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//writer.WriteHeader(http.StatusCreated)
}

func (h *Handler) deleteBook(c *gin.Context) {
	//idRequest := mux.Vars(request)
	//id := idRequest["id"]
	//err := a.Data.Delete(id)
	//logrus.Println(id)
	//if err != nil {
	//	log.Println("book hasn't been deleted(")
	//	writer.WriteHeader(http.StatusBadRequest)
	//	return
	//}
	//err = json.NewEncoder(writer).Encode(id)
	//if err != nil {
	//	log.Println(err)
	//	writer.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//writer.WriteHeader(http.StatusCreated)
}
