package services

import (
	"context"
	"teach/internal/httpclient"
	"teach/model"
	"teach/pkg/connector"

	storage "teach/storage"

	"go.mongodb.org/mongo-driver/mongo"
)

type ConfigService struct {
	ConfigStorage *storage.ConfigStorage
	httpclient    *httpclient.HTTPClient
}

type IConfigService interface {
	GetUserRedisService(ctx context.Context, userInformation model.UserInformationRequest) error
}

func NewConfigService(
	r *connector.Redis,
	db *mongo.Database,
	ht *httpclient.HTTPClient,
) IConfigService {
	return &ConfigService{
		ConfigStorage: storage.NewConfigStorage(db),
	}
}
