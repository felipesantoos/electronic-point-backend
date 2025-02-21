package redis

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/infra"
	"eletronic_point/src/utils"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var logger = infra.Logger().With().Str("port", "redis").Logger()

func valueExists(conn *redis.Client, key, value string) (bool, errors.Error) {
	storedValue, err := getValueFromKey(conn, key)
	if err != nil {
		return false, err
	}
	return storedValue == value, nil
}

func getRedisAddress() string {
	return fmt.Sprintf("%s:%s", utils.GetenvWithDefault("REDIS_HOST", "redis"), utils.GetenvWithDefault("REDIS_PORT", "6379"))
}

func getConnection() (*redis.Client, errors.Error) {
	conn := redis.NewClient(&redis.Options{
		Addr:     getRedisAddress(),
		Password: utils.GetenvWithDefault("REDIS_PASSWORD", ""),
		DB:       0,
	})
	if result := conn.Ping(); result.Err() != nil {
		logger.Log().Msg(fmt.Sprintf("an error occurred when trying to connect to the redis instance: %s", result.Err().Error()))
		return nil, errors.NewUnexpected()
	}
	return conn, nil
}

func getValueFromKey(conn *redis.Client, key string) (string, errors.Error) {
	result, err := conn.Get(key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		return "", errors.NewUnexpected()
	}
	return result, nil
}

func getKeyDuration(conn *redis.Client, key string) (*time.Duration, errors.Error) {
	result, err := conn.TTL(key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, errors.NewUnexpected()
	}
	return &result, nil
}
