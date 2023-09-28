package model

import (
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

type Post struct {
	BaseModel
	Title   string         `json:"title" gorm:"column:title;type:varchar(250);not null"`
	Content string         `json:"content" gorm:"column:content;type:varchar;"`
	Slug    string         `json:"slug" gorm:"column:slug;type:varchar(50);not null"`
	Image   string         `json:"image" gorm:"column:image;type:varchar(250);"`
	UserId  uuid.UUID      `json:"user_id" gorm:"column:user_id;type:uuid"`
	Files   pq.StringArray `json:"files" gorm:"column:files;type:varchar(500)[]" swaggertype:"array,string"`
	User    *User          `json:"user,omitempty" gorm:"foreignKey:UserId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Posts []Post

func (Post) TableName() string {
	return "posts"
}
