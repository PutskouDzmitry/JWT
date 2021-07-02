package repository

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"time"
)

func NewRedisDb(host string, port string, password string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", host, port),
		Password: password,
		DB:       0,
	})
	back := config()
	for {
		timeWait := back.NextBackOff()
		time.Sleep(timeWait)
		_, err := client.Ping().Result()
		if err != nil {
			logrus.Error("we wait connect to redis, time: ", timeWait)
		} else {
			break
		}
	}
	return client, nil
}
