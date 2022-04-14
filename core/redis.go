package core

import (
	"os"

	"github.com/go-redis/redis"
)

func RedisConnect() (*redis.Client, error) {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	pass := os.Getenv("REDIS_PASSWORD")
	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: pass,
		DB:       0,
	})
	return client, Ping(client)
}

func Ping(client *redis.Client) error {
	_, err := client.Ping().Result()
	return err
}
