package storage

import (
	"encoding/json"
	"fmt"
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
	err := rs.Redis.Set("information_"+userID, info.ToJson(), 3600)
	if err != nil {
		return err
	}
	return nil

}

func (rs RedisStorage) GetUserInformation(userID string) (*model.InitInformationRedis, error) {
	cmd := rs.Redis.Get("information_" + userID)
	var info model.InitInformationRedis
	val, err := cmd.Result()
	if err != nil {
		return nil, nil
	}
	err = json.Unmarshal([]byte(val), &info)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON data: %v", err)
	}
	return &info, nil
}
