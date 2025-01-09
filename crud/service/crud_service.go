package service

import (
	"context"
	"fmt"

	"teach/model"
)

func (cs CrudService) ReadAllData(ctx context.Context) ([]*model.UserData, error) {

	data, err := cs.storage.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println("data = ", data)

	return data, nil
}

func (cs CrudService) ReadDataId(ctx context.Context, id int) (*model.UserData, error) {

	data, err := cs.storage.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	fmt.Println("data = ", data)

	return data, nil
}

func (cs CrudService) CreateData(ctx context.Context, data model.UserData) error {

	err := cs.storage.InsertData(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (cs CrudService) UpdateData(ctx context.Context, id int, age int) error {
	isFind, err := cs.storage.IsFind(ctx, id)
	if err != nil {
		return err
	}

	if !isFind {
		return fmt.Errorf("%v not found", id)
	} else {
		err := cs.storage.UpdateAgeData(ctx, id, age)
		if err != nil {
			return err
		}
		fmt.Println(fmt.Sprintf("Updated %v's age to %v", id, age))
	}

	return nil
}

func (cs CrudService) DeleteData(c context.Context, id int) error {
	isFind, err := cs.storage.IsFind(c, id)
	if err != nil {
		return err
	}
	if isFind {
		err = cs.storage.DeleteDataId(c, id)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("%v not found", id)
	}

	return nil
}
