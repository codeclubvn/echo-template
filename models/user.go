package models

type User struct {
	BaseModel
	UserName string `gorm:"type:varchar(100);not null" json:"user_name"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
	Email    string `gorm:"type:varchar(100);not null" json:"email"`
	Avatar   string `gorm:"type:varchar(255);" json:"avatar"`
	Social   string `gorm:"type:varchar(255);" json:"social"`
	SocialId string `gorm:"type:varchar(255);" json:"social_id"`
	Post     []Post
}

func (User) TableName() string {
	return "users"
}
