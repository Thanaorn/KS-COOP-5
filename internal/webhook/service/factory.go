package service

import (
	"teach/storage"

	"go.mongodb.org/mongo-driver/mongo"
)

type WebhookService struct {
	storage *storage.Storage
}

type IWebhookService interface {
	WebhookService() error
}

func NewWebhookService(
	db *mongo.Database,
) IWebhookService {
	NewStorage := storage.NewStorage(db)
	return WebhookService{
		storage: NewStorage,
	}
}
