package repository

import (
	"context"
	"fmt"
	"github.com/PutskouDzmitry/golang-training-Library/pkg/entity"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type AuthMongo struct {
	mongo *mongo.Client
	redis *redis.Client
}

func (a *AuthMongo) SetAccessToken(token string, id string) error {
	fmt.Println(token, id)
	err := a.redis.Set(fmt.Sprintf("Access_token-%v", id), token, time.Minute*30).Err()
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthMongo) GetAccessToken(token string, id string) (string, error) {
	val, err := a.redis.Get(fmt.Sprintf("Access_token-%v", id)).Result()
	if err != nil {
		return "", err
	}
	if val != token {
		return "", fmt.Errorf("your token doesn't equal to original token")
	}
	return token, err
}

func (a *AuthMongo) SetRefreshToken(token string, id string) error {
	err := a.redis.Set(fmt.Sprintf("Refresh_token-%v", id), token, time.Hour*12).Err()
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthMongo) GetRefreshToken(token string, id string) (string, error) {
	val, err := a.redis.Get(fmt.Sprintf("Refresh_token-%v", id)).Result()
	if err != nil {
		return "", err
	}
	if val != token {
		return "", fmt.Errorf("your token doesn't equal to original token")
	}
	return token, err
}

func NewAuthMongo(mongo *mongo.Client, redis *redis.Client) *AuthMongo {
	return &AuthMongo{mongo: mongo, redis: redis}
}

type userMongo struct {
	Id       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
}

func (a *AuthMongo) CreateUser(user entity.User) (string, error) {
	db := a.mongo.Database("book")
	collection := db.Collection("login")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//db.Drop(ctx)
	idObj, err := primitive.ObjectIDFromHex(user.Id)
	if err != nil {
		return "", err
	}
	userM := userMongo{
		Id:       idObj,
		Username: user.Username,
		Password: user.Password,
	}
	result, err := collection.InsertOne(ctx, userM)
	if err != nil {
		logrus.Error("err ", err)
		return "", err
	}
	if result == nil {
		logrus.Error("result ", err)
		return "", err
	}
	return userM.Id.String(), nil
}

func (a *AuthMongo) GetUser(id primitive.ObjectID, username string, password string) (entity.User, error) {
	db := a.mongo.Database("book")
	collection := db.Collection("login")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var user entity.User
	filter := bson.M{}
	cur := collection.FindOne(ctx, filter)
	if err := cur.Decode(&user); err != nil {
		return entity.User{}, fmt.Errorf("error after decode %v", err)
	}
	return user, nil
}
