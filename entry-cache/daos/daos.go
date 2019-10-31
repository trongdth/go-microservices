package daos

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"github.com/trongdth/go_microservices/entry-cache/config"
)

var (
	redisClient *redis.Client
)

// Init : init redis
func Init(conf *config.Config) error {
	var err error
	redisClient := redis.NewClient(&redis.Options{
		Addr:     conf.Redis,
		Password: conf.RedisPwd, // no password set
		DB:       0,             // use default DB
	})

	pong, err := redisClient.Ping().Result()
	if pong == "PONG" {
		fmt.Println("Redis connect successfully!")
	} else {
		return errors.Wrap(err, "Redis.Init")
	}

	return nil
}

// GetRedisClient : return redis client
func GetRedisClient() *redis.Client {
	return redisClient
}
