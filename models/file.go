package models

import "encoding/json"

type File struct {
	BaseModel
	FileName      string          `json:"file_name" gorm:"column:file_name;type:varchar(50);not null"`
	Path          string          `json:"path" gorm:"column:path;type:varchar(255);not null"`
	Size          int64           `json:"size" gorm:"column:size;type:bigint;not null"`
	ExtensionName string          `json:"type" gorm:"column:extension_name;type:varchar(10);not null"`
	Data          json.RawMessage `json:"data" gorm:"column:data;type:jsonb;"` // save data flexibly
}

func (File) TableName() string {
	return "files"
}
