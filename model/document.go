package model

import (
	"gorm.io/gorm"
)

type Document struct {
	gorm.Model
	Name           string                 `gorm:"column:name"`
	SourceFileId   string                 `gorm:"column:source_file_id"`
	PreviewFileId  string                 `gorm:"column:preview_file_id"`
	EditableFileId string                 `gorm:"column:editable_file_id"`
	Metadata       map[string]interface{} `gorm:"column:metadata"`
	CreatedBy      string                 `gorm:"column:created_by"`
}
