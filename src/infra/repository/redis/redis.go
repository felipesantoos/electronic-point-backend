package redis

import (
	"dit_backend/src/core/utils"
	"dit_backend/src/infra"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var errLogger = Logger().Level(zerolog.ErrorLevel)

func valueExists(conn *redis.Client, key, value string) (bool, infra.Error) {
	storedValue, err := getValueFromKey(conn, key)
	if err != nil {
		return false, err
	}
	return storedValue == value, nil
}

func Logger() zerolog.Logger {
	return log.With().Str("layer", "infra|redis").Logger()
}

func getRedisAddress() string {
	return fmt.Sprintf("%s:%s", utils.GetenvWithDefault("REDIS_HOST", "redis"), utils.GetenvWithDefault("REDIS_PORT", "6379"))
}

func getConnection() (*redis.Client, infra.Error) {
	conn := redis.NewClient(&redis.Options{
		Addr:     getRedisAddress(),
		Password: utils.GetenvWithDefault("REDIS_PASSWORD", ""),
		DB:       0,
	})
	if result := conn.Ping(); result.Err() != nil {
		errLogger.Log().Msg(fmt.Sprintf("an error occurred when trying to connect to the redis instance: %s", result.Err().Error()))
		return nil, infra.NewUnexpectedSourceErr()
	}
	return conn, nil
}

func getValueFromKey(conn *redis.Client, key string) (string, infra.Error) {
	result, err := conn.Get(key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		return "", infra.NewUnexpectedSourceErr()
	}
	return result, nil
}
