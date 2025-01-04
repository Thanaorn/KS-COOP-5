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
	// CrudService() error
	// createData(c echo.Context) error
	ReadAllData(context.Context) ([]*model.TestData, error)
	// updateData(c echo.Context) error
	// deleteData(c echo.Context) error
}

func NewCrudService(db *mongo.Database) ICrudService {
	NewStorage := storage.NewStorage(db)
	return CrudService{
		storage: NewStorage,
	}
}
