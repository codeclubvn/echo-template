package request

type RegisterRequest struct {
	Email    string `json:"email" binding:"required" validate:"email"`
	Password string `json:"password" binding:"required" validate:"min=6,max=20"`
}

type UserGoogleRequest struct {
	Email     string `json:"email" binding:"required" validate:"email"`
	GoogleID  string `json:"id" binding:"required"`
	FirstName string `json:"family_name"`
	LastName  string `json:"given_name"`
}
