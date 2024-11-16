package model

import "time"

type File struct {
	ID           string    `gorm:"primaryKey,column:id"`
	Name         string    `gorm:"column:name;uniqueIndex"`
	OriginalName string    `gorm:"column:original_name"`
	MimeType     string    `gorm:"column:mimetype"`
	Size         int64     `gorm:"column:size"`
	Path         string    `gorm:"column:path;uniqueIndex"`
	Extension    string    `gorm:"column:extension"`
	Type         string    `gorm:"column:type"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}
