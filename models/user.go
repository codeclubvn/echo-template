package models

type User struct {
	BaseModel
	Name     string `gorm:"type:varchar(255);not null" json:"name"`
	UserName string `gorm:"type:varchar(100);not null" json:"user_name"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
	Email    string `gorm:"type:varchar(100);not null" json:"email"`
	Avatar   string `gorm:"type:varchar(255);" json:"avatar"`
	Social   string `gorm:"type:varchar(255);" json:"social"`
	SocialId string `gorm:"type:varchar(255);" json:"social_id"`
	Post     []Post `gorm:"foreignKey:UserId" json:"post"`
}

func (User) TableName() string {
	return "users"
}
