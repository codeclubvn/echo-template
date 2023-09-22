package models

import uuid "github.com/satori/go.uuid"

type Post struct {
	BaseModel
	Title   string    `json:"title" gorm:"column:title;type:varchar(50);not null"`
	Content string    `json:"content" gorm:"column:content;type:varchar;"`
	Slug    string    `json:"slug" gorm:"column:slug;type:varchar(50);not null"`
	Image   string    `json:"image" gorm:"column:image;type:varchar(50);"`
	UserId  uuid.UUID `json:"user_id" gorm:"column:user_id;type:uuid"`
	User    User      `json:"user" gorm:"foreignKey:UserId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	File    []File
}

func (Post) TableName() string {
	return "posts"
}
