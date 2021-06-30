package service

import (
	"github.com/PutskouDzmitry/golang-training-Library/pkg/repository"
	"github.com/PutskouDzmitry/golang-training-Library/pkg/struct"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Authorization interface {
	CreateUser(user _struct.User) (string, error)
	GenerateToken(id primitive.ObjectID,username string, password string) (string, error)
	ParseToken(token string) (string, error)
}

type Books interface {

}

type Service struct {
	Authorization
	Books
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
