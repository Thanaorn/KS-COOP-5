package model

import (
	"encoding/json"
)

type AddressRedis struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

type ContactRedis struct {
	Email   string       `json:"email"`
	Phone   string       `json:"phone"`
	Address AddressRedis `json:"address"`
}

type InitInformationRedis struct {
	UserID  string       `json:"user_id"`
	Name    string       `json:"name"`
	Age     string       `json:"age"`
	Contact ContactRedis `json:"contact"`
}

func (i InitInformationRedis) ToJson() string {
	jsonData, _ := json.Marshal(i)
	return string(jsonData)
}
