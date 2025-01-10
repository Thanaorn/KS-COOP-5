package model

type StatusResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type AddressResponse struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

type ContactResponse struct {
	Email   string          `json:"email"`
	Phone   string          `json:"phone"`
	Address AddressResponse `json:"address"`
}

type UserInformationResponse struct {
	UserID  string          `json:"user_id"`
	Name    string          `json:"name"`
	Age     string          `json:"age"`
	Contact ContactResponse `json:"contact"`
}
