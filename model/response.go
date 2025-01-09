package model

type StatusResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type UserInformationRespond struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Age    string `json:"age"`
}
