package models

import uuid "github.com/satori/go.uuid"

type File struct {
	BaseModel
	Name   string    `json:"name" gorm:"column:name;type:varchar(50);not null"`
	Path   string    `json:"path" gorm:"column:path;type:varchar(50);not null"`
	Data   string    `json:"data" gorm:"column:data;type:varchar(50);not null"`
	PostId uuid.UUID `json:"post_id" gorm:"column:post_id;type:uuid"`
}

func (File) TableName() string {
	return "files"
}
