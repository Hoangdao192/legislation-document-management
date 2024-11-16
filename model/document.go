package model

import (
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
)

type Document struct {
	gorm.Model
	Name           string   `gorm:"column:name"`
	SourceFileId   string   `gorm:"column:source_file_id"`
	PreviewFileId  string   `gorm:"column:preview_file_id"`
	EditableFileId string   `gorm:"column:editable_file_id"`
	Metadata       Metadata `gorm:"column:metadata;type:json"`
	CreatedBy      string   `gorm:"column:created_by"`
}

type Metadata map[string]interface{}

func (m Metadata) Value() (driver.Value, error) {
	if len(m) == 0 {
		return nil, nil
	}
	return json.Marshal(m)
}

func (m *Metadata) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), m)
}
