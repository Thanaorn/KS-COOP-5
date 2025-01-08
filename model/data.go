package model

type UserData struct {
	Id   int    `json:"user_id" bson:"user_id"`
	Name string `json:"user_name" bson:"user_name"`
	Age  int    `json:"user_age" bson:"user_age"`
}
