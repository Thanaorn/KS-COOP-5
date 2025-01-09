package model

import "encoding/json"

type InitInformationRedis struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Age    string `json:"age"`
}

func (i InitInformationRedis) ToJson() string {
	jsonData, _ := json.Marshal(i)
	return string(jsonData)
}
