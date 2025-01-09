package services

import (
	"context"
	"teach/model"
)

func (cs ConfigService) GetUserRedisService(ctx context.Context, userInformation model.UserInformationRequest) error {
	err := cs.httpclient.RedisStorage.SetUserInformation(userInformation.UserID, model.InitInformationRedis{
		UserID: userInformation.UserID,
		Name:   userInformation.Name,
		Age:    userInformation.Age,
	})
	if err != nil {
		return err
	}
	return nil
}
