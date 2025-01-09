package model

type StatusResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
