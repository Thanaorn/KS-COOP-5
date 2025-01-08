package connector

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	Client  *redis.Client
	Context context.Context
}

func NewRedisClient(redisURI, address, password string) *Redis {
	var err error = nil
	var redisConfig *redis.Options = &redis.Options{}
	if redisURI != "" {
		redisConfig, err = redis.ParseURL(redisURI)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		redisConfig = &redis.Options{
			Addr:     address,
			Password: password,
			DB:       0,
		}
	}
	rdb := redis.NewClient(redisConfig)
	return &Redis{
		Client:  rdb,
		Context: context.Background(),
	}
}

func (r Redis) Disconnect() error {
	return r.Client.Close()
}

func (r Redis) Ping() *redis.StatusCmd {
	return r.Client.Ping(r.Context)
}

func (r Redis) Set(key string, value interface{}, duration int64) error {
	statusCmd := r.Client.Set(r.Context, key, value, time.Duration(duration)*time.Second)
	if statusCmd.Err() != nil {
		return statusCmd.Err()
	}
	return nil
}

func (r Redis) Get(key string) *redis.StringCmd {
	return r.Client.Get(r.Context, key)
}

func (r *Redis) GetString(key string) (string, error) {
	cmd := r.Client.Get(r.Context, key)
	val, err := cmd.Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *Redis) GetInt(key string) (int, error) {
	cmd := r.Client.Get(r.Context, key)
	val, err := cmd.Result()
	if err != nil {
		return 0, err
	}

	// Convert the value from string to int
	intVal, convErr := strconv.Atoi(val)
	if convErr != nil {
		return 0, convErr
	}

	return intVal, nil
}

func (r *Redis) Clear(key string) error {
	statusCmd := r.Client.Del(r.Context, key)
	if statusCmd.Err() != nil {
		return statusCmd.Err()
	}
	return nil
}

func (r *Redis) Exists(key string) (int64, error) {
	exists, err := r.Client.Exists(r.Context, key).Result()
	if err != nil {
		return 0, err
	}
	return exists, nil
}
