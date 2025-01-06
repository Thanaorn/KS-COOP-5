package service

import (
	"context"
	"teach/model"
	"teach/storage"

	"go.mongodb.org/mongo-driver/mongo"
)

type CrudService struct {
	storage *storage.Storage
}

type ICrudService interface {
	CreateData(ctx context.Context, data model.TestData) error
	ReadAllData(ctx context.Context) ([]*model.TestData, error)
	UpdateData(ctx context.Context, name string, age int) error
	// deleteData(c echo.Context) error
}

func NewCrudService(db *mongo.Database) ICrudService {
	NewStorage := storage.NewStorage(db)
	return CrudService{
		storage: NewStorage,
	}
}
