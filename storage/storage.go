package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type data struct {
	UserName string `json:"name"`
	Age      int    `json:"age"`
}

type Storage struct {
	storage *mongo.Collection
}

type ConfigStorage struct {
	ConfigCollection *mongo.Collection
}

func NewStorage(db *mongo.Database) *Storage {
	return &Storage{
		storage: db.Collection(os.Getenv("MONGO_COLLECTION")),
	}
}
func NewConfigStorage(db *mongo.Database) *ConfigStorage {
	return &ConfigStorage{
		ConfigCollection: db.Collection(os.Getenv("MONGO_COLLECTION")),
	}
}

func (s Storage) FindAll(ctx context.Context) (*data, error) {
	filter := bson.D{{}}
	cursor, err := s.storage.Find(ctx, filter)
	if err == mongo.ErrNoDocuments {
		fmt.Println(err)
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
	var data data
	if err := json.Unmarshal(jsonData, &data); err != nil {
		log.Fatalf("Error unmarshalling JSON data: %v", err)
	}
	return &data, nil
}
