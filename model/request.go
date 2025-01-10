package model

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

type Contact struct {
	Email   string  `json:"email"`
	Phone   string  `json:"phone"`
	Address Address `json:"address"`
}

type UserInformationRequest struct {
	UserID  string  `json:"user_id"`
	Name    string  `json:"name"`
	Age     string  `json:"age"`
	IDCard  string  `json:"id_card"`
	Contact Contact `json:"contact"`
}

type UserIDInformationRequest struct {
	UserId string `json:"user_id"`
}
