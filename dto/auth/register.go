package dto

type RegisterRequest struct {
	Email       string `json:"email" binding:"required" validate:"email"`
	Password    string `json:"password" binding:"required" validate:"min=6,max=20"`
	RequestFrom string `json:"request_from" binding:"required" enums:"erp/,web,app"`
}

type UserGoogleRequest struct {
	Email     string `json:"email" binding:"required" validate:"email"`
	GoogleID  string `json:"id" binding:"required"`
	FirstName string `json:"family_name"`
	LastName  string `json:"given_name"`
}
