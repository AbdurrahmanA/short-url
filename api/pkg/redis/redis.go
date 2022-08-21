package redis

import (
	"context"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

type Redis struct {
	RDB *redis.Client
	CTX context.Context
}

var redisClient *redis.Client
var doOnce sync.Once

func NewRedisClient(URL string, Password string) *Redis {
	context := context.Background()

	doOnce.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     URL,
			Password: Password,
		})

		_, err := redisClient.Ping().Result()
		if err != nil {
			zap.S().Error("Error while connecting to Redis", err)
		}
	})
	zap.S().Info("Redis connected successfully")
	return &Redis{
		RDB: redisClient,
		CTX: context,
	}
}

func (r *Redis) Ping() error {
	return r.RDB.Ping().Err()
}

func (r *Redis) Set(key string, value string, ttl time.Duration) error {
	return r.RDB.Set(key, value, ttl).Err()
}

func (r *Redis) Get(key string) (string, error) {
	return r.RDB.Get(key).Result()
}
