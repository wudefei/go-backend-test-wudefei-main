package redis

import (
	"user-api/internal/config"

	"github.com/gomodule/redigo/redis"
)

var redisClient *redis.Pool

func InitRedis() {
	cfg, ok := config.Config.Redis["default"]
	if !ok {
		panic("redis config not exit")
	}

	redisClient = &redis.Pool{
		MaxIdle:   cfg.MaxIdle,
		MaxActive: cfg.MaxActive,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", cfg.Uri,
				redis.DialPassword(cfg.Auth),
				redis.DialDatabase(cfg.Db),
			)
			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}
}

func GetRedis() *redis.Pool {
	return redisClient
}
