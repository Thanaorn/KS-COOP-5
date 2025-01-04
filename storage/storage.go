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
