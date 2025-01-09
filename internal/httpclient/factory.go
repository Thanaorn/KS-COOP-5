package httpclient

import (
	"net/http"
	"teach/pkg/connector"
	storage "teach/storage"

	"go.mongodb.org/mongo-driver/mongo"
)

type HTTPClient struct {
	Client        http.Client
	ConfigStorage *storage.ConfigStorage
	RedisStorage  *storage.RedisStorage
}

func NewHTTPClient(c http.Client, db *mongo.Database, rd *connector.Redis) *HTTPClient {
	ConfigStorage := storage.NewConfigStorage(db)
	Redis := storage.NewRedisStorage(rd)
	return &HTTPClient{
		Client:        c,
		ConfigStorage: ConfigStorage,
		RedisStorage:  Redis,
	}
}
