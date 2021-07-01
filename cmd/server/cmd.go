package main

import (
	"fmt"
	"github.com/PutskouDzmitry/golang-training-Library/pkg/handler"
	"github.com/PutskouDzmitry/golang-training-Library/pkg/repository"
	"github.com/PutskouDzmitry/golang-training-Library/pkg/server"
	"github.com/PutskouDzmitry/golang-training-Library/pkg/service"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

var (
	user     = os.Getenv("DB_USERS_USER")
	password = os.Getenv("DB_USERS_PASSWORD")
	host = os.Getenv("DB_USERS_HOST")
	port = os.Getenv("DB_USER_PORT")
)

func initValues() {
	if user == "" {
		user = "root"
	}
	if password == "" {
		password = "example"
	}
	if host == "" {
		host = "localhost"
		//host = "mongo"
	}
	if port == "" {
		port = "27017"
	}
}

func initClient(user string, password string, host string, port string) string{
	return fmt.Sprintf("mongodb://%v:%v/?sslmode=disable", host, port)
}

//func initClient(user string, password string, host string, port string) string{
//	return fmt.Sprintf("mongodb://%v:%v@%v:%v/?sslmode=disable", user, password, host, port)
//}

func main() {
	initValues()
	mongo, err := repository.NewMongodb(user, password, host, port)
	if err != nil {
		logrus.Fatal(err)
	}
	redis, err := repository.NewRedisDb()
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

