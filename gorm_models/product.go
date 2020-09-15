package gorm_models

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID          uint `gorm:"primarykey"`
	Name        string
	Description string
	ImageSrc    string
	Price       float64
	CategoryID  uint
	Category    Category `gorm:"joinForeignKey:CategoryID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
