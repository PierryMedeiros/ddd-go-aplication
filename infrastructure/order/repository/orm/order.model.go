package orm

import (
    "gorm.io/gorm"
    "desafio-ddd-go/infrastructure/customer/repository/orm"
    "desafio-ddd-go/infrastructure/orderitem"
)

type OrderModel struct {
    ID         string          `gorm:"primaryKey"`
    CustomerID string          `gorm:"not null"`
    Customer   orm.CustomerModel   `gorm:"foreignKey:CustomerID;references:ID"`
    Items      []orderitem.OrderItemModel `gorm:"foreignKey:OrderID"`
    Total      float64         `gorm:"not null"`
}

func Migrate(db *gorm.DB) error {
    return db.AutoMigrate(&OrderModel{})
}
