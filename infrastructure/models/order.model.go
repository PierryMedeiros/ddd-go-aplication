package models

import (
    "gorm.io/gorm"

)

type OrderModel struct {
    ID         string          `gorm:"primaryKey"`
    CustomerID string          `gorm:"not null"`
    Customer   CustomerModel   `gorm:"foreignKey:CustomerID;references:ID"`
    Items      []OrderItemModel `gorm:"foreignKey:OrderID"`
    Total      float64         `gorm:"not null"`
}

func (*OrderModel) Migrate(db *gorm.DB) error {
    return db.AutoMigrate(&OrderModel{})
}
