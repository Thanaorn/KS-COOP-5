package service

import (
	"context"
	"fmt"

	"teach/model"
)

func (cs CrudService) ReadAllData(ctx context.Context) ([]*model.TestData, error) {

	data, err := cs.storage.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println("data = ", data)

	return data, nil
}

func (cs CrudService) CreateData(ctx context.Context, data model.TestData) error {

	err := cs.storage.InsertData(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (cs CrudService) UpdateData(ctx context.Context, name string, age int) error {
	isFind, err := cs.storage.FindByName(ctx, name)
	if err != nil {
		return err
	}

	if !isFind {
		return fmt.Errorf("%s not found", name)
	} else {
		err := cs.storage.UpdateAgeData(ctx, name, age)
		if err != nil {
			return err
		}
		fmt.Println(fmt.Sprintf("Updated %s's age to %d", name, age))
	}

	return nil
}
