package gorm_models

import (
	"gorm.io/gorm"
	"time"
)

// Product
type Product struct {
	ID          uint `gorm:"primarykey"`
	Name        string
	Description string
	ImageSrc    string
	Price       float64
	CategoryID  uint
	Category    Category               `gorm:"joinForeignKey:CategoryID"`
	Properties  []ProductPropertyValue `gorm:"foreignKey:ProductID;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

// Product property
type ProductProperty struct {
	Code      string `gorm:"primarykey"`
	Name      string
	Required  bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (ProductProperty) TableName() string {
	return "product_properties"
}

// Product property values
type ProductPropertyValue struct {
	ID           uint `gorm:"primarykey"`
	ProductID    uint
	PropertyCode string
	Property     ProductProperty `gorm:"foreignKey:PropertyCode;references:Code"`
	Value        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (ProductPropertyValue) TableName() string {
	return "product_property_values"
}
