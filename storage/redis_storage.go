package storage

import (
	// "encoding/json"
	// "fmt"
	"teach/model"
	"teach/pkg/connector"
)

type RedisStorage struct {
	Redis *connector.Redis
}

func NewRedisStorage(redis *connector.Redis) *RedisStorage {
	return &RedisStorage{
		Redis: redis,
	}
}

func (rs RedisStorage) SetUserInformation(userID string, info model.InitInformationRedis) error {
	//Implement

	return nil

}

func (rs RedisStorage) GetUserInformation(userID string) (*model.InitInformationRedis, error) {
	//Implement

	return nil, nil
}

func (rs RedisStorage) ClearUserInformation(userID string) error {
	//Implement

	return nil

}
