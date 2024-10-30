package orderitem

import (
    "gorm.io/gorm"
    ormr"desafio-ddd-go/infrastructure/product/repository/orm"
    orms "desafio-ddd-go/infrastructure/order/repository/orm"
)

type OrderItemModel struct {
    ID        string      `gorm:"primaryKey"`
    ProductID string      `gorm:"not null"`
    Product   ormr.ProductModel `gorm:"foreignKey:ProductID;references:ID"`
    OrderID   string      `gorm:"not null"`
    Order     orms.OrderModel  `gorm:"foreignKey:OrderID;references:ID"`
    Quantity  int         `gorm:"not null"`
    Name      string      `gorm:"not null"`
    Price     float64     `gorm:"not null"`
}

func Migrate(db *gorm.DB) error {
    return db.AutoMigrate(&OrderItemModel{})
}
