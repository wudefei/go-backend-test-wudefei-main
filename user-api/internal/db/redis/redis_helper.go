package redis

import (
	"github.com/gomodule/redigo/redis"
)

type RedisHelper struct {
	LogId string
}

// NewRedisHelper 获取redisHelper
func NewRedisHelper(logId string) *RedisHelper {
	return &RedisHelper{LogId: logId}
}

// StringWithExpire redis写string 数据
func (r *RedisHelper) StringWithExpire(key, value string, expire int) error {
	cli := redisClient.Get()
	defer cli.Close()

	_, err := cli.Do("SETEX", key, expire, value)
	return err
}

//StringGet 获取redis string数据
func (r *RedisHelper) StringGet(key string) (string, error) {
	cli := redisClient.Get()
	defer cli.Close()

	ret, err := redis.String(cli.Do("GET", key))
	if err != nil && err == redis.ErrNil { //还没有key
		return "", nil
	}

	return ret, err
}

//SET key value [EX seconds] [PX millisecounds] [NX|XX]
func (r *RedisHelper) SetNxWithMillSecExpire(key, value string, millSecExpire int) error {
	cli := redisClient.Get()
	defer cli.Close()

	_, err := cli.Do("SET", key, value, "PX", millSecExpire, "NX")
	return err
}
