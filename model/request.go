package model

type UserInformationRequest struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
}
