package redis

import (
	"backend_template/src/core"
	"backend_template/src/core/domain/errors"
	"backend_template/src/core/utils"
	"backend_template/src/infra"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/rs/zerolog"
)

var logger = infra.Logger().With().Str("port", "redis").Logger()

func valueExists(conn *redis.Client, key, value string) (bool, errors.Error) {
	storedValue, err := getValueFromKey(conn, key)
	if err != nil {
		return false, err
	}
	return storedValue == value, nil
}

func Logger() zerolog.Logger {
	return core.CoreLogger().With().Str("layer", "infra|redis").Logger()
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
