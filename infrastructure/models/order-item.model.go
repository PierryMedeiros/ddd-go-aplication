package models

import (
    "gorm.io/gorm"
    
)

type OrderItemModel struct {
    ID        string      `gorm:"primaryKey"`
    ProductID string      `gorm:"not null"`
    Product   ProductModel `gorm:"foreignKey:ProductID;references:ID"`
    OrderID   string      `gorm:"not null"`
    Order     OrderModel  `gorm:"foreignKey:OrderID;references:ID"`
    Quantity  int         `gorm:"not null"`
    Name      string      `gorm:"not null"`
    Price     float64     `gorm:"not null"`
}

func (*OrderItemModel) Migrate(db *gorm.DB) error {
    return db.AutoMigrate(&OrderItemModel{})
}
