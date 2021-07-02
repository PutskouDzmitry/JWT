package repository

import (
	"github.com/PutskouDzmitry/golang-training-Library/pkg/entity"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Authorization interface {
	CreateUser(user entity.User) (string, error)
	GetUser(id primitive.ObjectID, username, password string) (entity.User, error)
	SetAccessToken(token string, id string) error
	GetAccessToken(token string, id string) (string, error)
	SetRefreshToken(token string, id string) error
	GetRefreshToken(token string, id string) (string, error)
}

type BooksRepo interface {
	ReadAll() ([]entity.Book, error)
	Read(id string) (entity.Book, error)
	Add(book entity.Book) (string, error)
	Update(id string, value int) (string, error)
	Delete(id string) error
}

type Repository struct {
	Authorization
	BooksRepo
}

func NewRepository(mongo *mongo.Client, redis *redis.Client) *Repository {
	return &Repository{
		Authorization: NewAuthMongo(mongo, redis),
		BooksRepo:     NewBookData(mongo),
	}
}
