package services

import (
	"context"
	"teach/model"
)

func (cs ConfigService) GetUserRedisService(
	ctx context.Context,
	userID string) (
	*model.InitInformationRedis, error) {
	userInformation, err := cs.RedisStorage.GetUserInformation(userID)
	if err != nil {
		return nil, err
	}
	return userInformation, nil
}

func (cs ConfigService) SetUserRedisService(
	ctx context.Context,
	userID string,
	info model.InitInformationRedis) error {
	err := cs.RedisStorage.SetUserInformation(userID, info)
	if err != nil {
		return err
	}
	return nil
}

func (cs ConfigService) DeleteUserRedisService(
	ctx context.Context,
	userID string) error {
	err := cs.RedisStorage.ClearUserInformation(userID)
	if err != nil {
		return err
	}
	return nil
}
