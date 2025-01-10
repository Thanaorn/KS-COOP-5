package service

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type CrudService struct {
}

type ICrudService interface {
	// CreateData(ctx context.Context, data model.UserData) error
	// ReadAllData(ctx context.Context) ([]*model.UserData, error)
	// ReadDataId(ctx context.Context, id int) (*model.UserData, error)
	// UpdateData(ctx context.Context, id int, age int) error
	// DeleteData(c context.Context, id int) error
}

func NewCrudService(db *mongo.Database) ICrudService {
	return CrudService{}
}
