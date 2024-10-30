package orm

import (
    "gorm.io/gorm"
)

type ProductModel struct {
    ID    string  `gorm:"primaryKey"`        
    Name  string  `gorm:"not null"`          
    Price float64 `gorm:"not null"`          
}

func Migrate(db *gorm.DB) error {
    return db.AutoMigrate(&ProductModel{}) 
}
