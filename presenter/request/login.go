package request

type LoginRequest struct {
	Email       string `json:"email" binding:"required" validate:"required"`
	Password    string `json:"password" binding:"required" validate:"min=6,max=20"`
	RequestFrom string `json:"request_from" binding:"required" enums:"trail/,web,app"`
}

type LoginByGoogleRequest struct {
	Email    string `json:"email" binding:"required" validate:"email"`
	GoogleId string `json:"google_id" binding:"required"`
}
