package repository

import (
	"github.com/PutskouDzmitry/golang-training-Library/pkg/struct"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Authorization interface {
	CreateUser(user _struct.User) (string, error)
	GetUser(id primitive.ObjectID, username, password string) (_struct.User, error)
	SetAccessToken(token string) (string, error)
	GetAccessToken(token string) (string, error)
	SetRefreshToken(token string) (string, error)
	GetRefreshToken(token string) (string, error)
}

type BooksRepo interface {
	ReadAll() ([]_struct.Book, error)
	Read(id string) (_struct.Book, error)
	Add(book _struct.Book) (string, error)
	Update(id string, value int) (string, error)
	Delete(id string) error
}

type Repository struct {
	Authorization
	BooksRepo
}

func NewRepository(mongo *mongo.Client, redis *redis.Client) *Repository{
	return &Repository{
		Authorization: NewAuthMongo(mongo, redis),
		BooksRepo: NewBookData(mongo),
	}
}
