package main

import (
	"github.com/PutskouDzmitry/golang-training-Library/pkg/handler"
	"github.com/PutskouDzmitry/golang-training-Library/pkg/repository"
	"github.com/PutskouDzmitry/golang-training-Library/pkg/server"
	"github.com/PutskouDzmitry/golang-training-Library/pkg/service"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

var (
	passwordRedis = os.Getenv("REDIS_PASSWORD")
	portRedis     = os.Getenv("REDIS_PORT")
	hostRedis     = os.Getenv("REDIS_HOST")
	userMongo     = os.Getenv("MONGO_USER")
	passwordMongo = os.Getenv("MONGO_PASSWORD")
	hostMongo     = os.Getenv("MONGO_HOST")
	portMongo     = os.Getenv("MONGO_PORT")
)

func initValues() {
	if passwordRedis == "none" {
		passwordRedis = ""
	}
	if portRedis == "" {
		portRedis = "6379"
	}
	if hostRedis == "" {
		hostRedis = "localhost"
	}
	if userMongo == "" {
		userMongo = "root"
	}
	if passwordMongo == "" {
		passwordMongo = "example"
	}
	if hostMongo == "" {
		hostMongo = "localhost"
	}
	if portMongo == "" {
		portMongo = "27017"
	}
}

func main() {
	initValues()
	mongo, err := repository.NewMongodb(userMongo, passwordMongo, hostMongo, portMongo)
	if err != nil {
		logrus.Fatal(err)
	}
	redis, err := repository.NewRedisDb(hostRedis, portRedis, passwordRedis)
	if err != nil {
		logrus.Fatal(err)
	}
	repos := repository.NewRepository(mongo, redis)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(server.Server)
	if err := srv.Run("8081", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server %s", err.Error())
	}
}
