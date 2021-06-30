package repository

import (
	"github.com/go-redis/redis"
)

func NewRedisDb() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, err
}