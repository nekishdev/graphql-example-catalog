package gorm_models

import (
	"gorm.io/gorm"
	"time"
)

type Category struct {
	ID          uint `gorm:"primarykey"`
	Name        string
	Description string
	ImageSrc    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
