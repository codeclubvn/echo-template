package models

import uuid "github.com/satori/go.uuid"

type File struct {
	BaseModel
	FileName string    `json:"file_name" gorm:"column:file_name;type:varchar(50);not null"`
	Size     int64     `json:"size" gorm:"column:size;type:bigint;not null"`
	Data     string    `json:"data" gorm:"column:data;type:varchar(50);not null"`
	PostId   uuid.UUID `json:"post_id" gorm:"column:post_id;type:uuid"`
	Post     Post      `json:"post" gorm:"foreignKey:PostId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (File) TableName() string {
	return "files"
}
