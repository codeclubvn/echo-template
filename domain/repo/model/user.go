package model

type User struct {
	BaseModel
	UserName string `gorm:"type:varchar(100);not null" json:"user_name"`
	Password string `gorm:"type:varchar(255);not null" json:"-"`
	Email    string `gorm:"type:varchar(100);not null" json:"email"`
	Avatar   string `gorm:"type:varchar(255);" json:"avatar"`
	Social   string `gorm:"type:varchar(255);" json:"social,omitempty"`
	SocialId string `gorm:"type:varchar(255);" json:"social_id,omitempty"`
	Post     Posts  `json:"posts,omitempty"`
}

func (User) TableName() string {
	return "users"
}
