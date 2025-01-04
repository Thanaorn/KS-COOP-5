package service

import (
	"context"
	"fmt"

	"teach/model"
)

// func (cs CrudService) CrudService() error {
// 	data, err := cs.storage.FindAll(context.TODO())
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println("data = ", data)

// 	return nil
// }

func (cs CrudService) ReadAllData(ctx context.Context) ([]*model.TestData, error) {

	data, err := cs.storage.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println("data = ", data)

	return data, nil
}
