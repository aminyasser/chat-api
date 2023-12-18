package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func GetRedis() *redis.Client {
	if client == nil {
		initRedis()
	}
	return client
}

func initRedis() {
	ctx := context.Background()

	client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err.Error())
	}

}
