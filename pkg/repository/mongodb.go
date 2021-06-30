package repository

import (x``
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func initClient(user string, password string, host string, port string) string{
	return fmt.Sprintf("mongodb://%v:%v/?sslmode=disable", host, port)
}

//func initClient(user string, password string, host string, port string) string{
//	return fmt.Sprintf("mongodb://%v:%v@%v:%v/?sslmode=disable", user, password, host, port)
//}

func NewMongodb(user string, password string, host string, port string) (*mongo.Client, error) {
	fmt.Println(host, port)
	client, err := mongo.NewClient(options.Client().ApplyURI(initClient(user, password, host, port)))
	if err != nil {
		logrus.Fatal("error with client ", err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		logrus.Fatal("error with connect to db ", err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		logrus.Fatal(err)
	}
	return client, err
}