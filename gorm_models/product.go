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
	Category    Category          `gorm:"joinForeignKey:CategoryID"`
	Properties  []ProductProperty `gorm:"foreignKey:ProductID;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type ProductProperty struct {
	ID         uint `gorm:"primarykey"`
	ProductID  uint
	PropertyID uint
	Property   Property `gorm:"joinForeignKey:PropertyID"`
	Value      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (ProductProperty) TableName() string {
	return "product_properties"
}
