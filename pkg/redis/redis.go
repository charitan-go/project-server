package redispkg

import (
	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func connect() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // No password set
		DB:       0,  // Use default DB
		Protocol: 2,  // Connection protocol
	})
}

func SetupRedis() {
	connect()
}
