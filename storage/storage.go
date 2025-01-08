package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"teach/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Storage struct {
	storage *mongo.Collection
}

func NewStorage(db *mongo.Database) *Storage {
	return &Storage{
		storage: db.Collection(os.Getenv("MONGO_COLLECTION")),
	}
}

func (s Storage) FindById(ctx context.Context, id int) (*model.UserData, error) {
	var data model.UserData
	err := s.storage.FindOne(ctx, bson.D{{"user_id", id}}).Decode(&data)
	if err != nil {
		fmt.Println("DOEST FIND NAME = ", id)
		return nil, err
	}
	fmt.Println("FIND NAME = ", id)
	return &data, nil
}

func (s Storage) IsFind(ctx context.Context, id int) (bool, error) {
	var data model.UserData
	err := s.storage.FindOne(ctx, bson.D{{"user_id", id}}).Decode(&data)
	if err != nil {
		fmt.Println("DOEST FIND NAME = ", id)
		return false, err
	}
	fmt.Println("FIND NAME = ", id)
	return true, nil
}

func (s Storage) UpdateAgeData(ctx context.Context, id int, age int) error {
	filter := bson.D{{"user_id", id}}
	update := bson.D{{"$set", bson.D{{"user_age", age}}}}
	res, err := s.storage.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	fmt.Println("res = ", res)
	return nil
}

func (s Storage) InsertData(ctx context.Context, data model.UserData) error {
	jsonbody, err := json.Marshal(data)
	if err != nil {
		fmt.Println("err = ", err)
	}
	fmt.Println("jsonbody = ", string(jsonbody))
	result, err := s.storage.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	fmt.Println("result = ", result)
	return nil
}

func (s Storage) FindAll(ctx context.Context) ([]*model.UserData, error) {
	filter := bson.D{{}}
	fmt.Println("filter = ", filter)
	cursor, err := s.storage.Find(ctx, filter)
	if err != nil {
		fmt.Println("err = ", err)
	}
	var results []bson.M
	for cursor.Next(context.TODO()) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}

		results = append(results, result)
	}
	if err := cursor.Err(); err != nil {
		log.Fatalf("Cursor error: %v", err)
	}
	jsonData, err := json.Marshal(results)
	if err != nil {
		return nil, err
	}
	var data []*model.UserData
	fmt.Println("jsonData = ", string(jsonData))
	if err := json.Unmarshal(jsonData, &data); err != nil {
		log.Fatalf("Error unmarshalling JSON data: %v", err)
	}
	return data, nil
}

func (s Storage) DeleteDataId(ctx context.Context, id int) error {
	filter := bson.D{{"user_id", id}}
	res, err := s.storage.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	fmt.Println("res = ", res)
	return nil
}
