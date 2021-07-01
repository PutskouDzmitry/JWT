package service

import (
	"github.com/PutskouDzmitry/golang-training-Library/pkg/repository"
	"github.com/PutskouDzmitry/golang-training-Library/pkg/struct"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Authorization interface {
	CreateUser(user _struct.User) (string, error)
	GenerateTokenAccessToken(id primitive.ObjectID,username string, password string) (string, error)
	ParseAccessToken(token string) (string, error)
	GenerateTokenRefreshToken(id primitive.ObjectID,username string, password string) (string, error)
	ParseRefreshToken(token string) (string, error)
}

type Books interface {
	ReadAll() ([]_struct.Book, error)
	Read(id string) (_struct.Book, error)
	Add(book _struct.Book) (string, error)
	Update(id string, value int) (string, error)
	Delete(id string) error
}

type Service struct {
	Authorization
	Books
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Books: NewBookService(repos.BooksRepo),
	}
}

