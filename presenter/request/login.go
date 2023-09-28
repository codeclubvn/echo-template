package request

type LoginRequest struct {
	Email    string `json:"email" binding:"required" validate:"email" example:"hieuhoccode@gmail.com"`
	Password string `json:"password" binding:"required" validate:"min=6,max=20" example:"hieuhoccode"`
}

type LoginByGoogleRequest struct {
	Email    string `json:"email" binding:"required" validate:"email"`
	GoogleId string `json:"google_id" binding:"required"`
}
