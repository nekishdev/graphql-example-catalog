package gorm_models

import (
	"gorm.io/gorm"
	"time"
)

type Property struct {
	ID        uint `gorm:"primarykey"`
	Name      string
	Code      string
	Required  bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Property) TableName() string {
	return "properties"
}
