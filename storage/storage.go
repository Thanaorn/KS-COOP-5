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

func (s Storage) FindByName(ctx context.Context, name string) (bool, error) {
	var data model.TestData
	err := s.storage.FindOne(ctx, bson.D{{"namedata", name}}).Decode(&data)
	if err != nil {
		fmt.Println("DOEST FIND NAME = ", name)
		return false, err
	}
	fmt.Println("FIND NAME = ", name)
	return true, nil
}

func (s Storage) UpdateAgeData(ctx context.Context, name string, age int) error {
	filter := bson.D{{"namedata", name}}
	update := bson.D{{"$set", bson.D{{"agedata", age}}}}
	res, err := s.storage.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	fmt.Println("res = ", res)
	return nil
}

func (s Storage) InsertData(ctx context.Context, data model.TestData) error {
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

func (s Storage) FindAll(ctx context.Context) ([]*model.TestData, error) {
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
	var data []*model.TestData
	fmt.Println("jsonData = ", string(jsonData))
	if err := json.Unmarshal(jsonData, &data); err != nil {
		log.Fatalf("Error unmarshalling JSON data: %v", err)
	}
	return data, nil
}
