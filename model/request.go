package model

type UserInformationRequest struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Age    string `json:"age"`
}

type UserIDInformationRequest struct {
	UserId string `json:"user_id"`
}
