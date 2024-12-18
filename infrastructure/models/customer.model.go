package models

import (
    "gorm.io/gorm"
)

type CustomerModel struct {
    ID           string `gorm:"primaryKey"`
    Name         string `gorm:"not null"`
    Street       string `gorm:"not null"`
    Number       int    `gorm:"not null"`
    Zipcode      string `gorm:"not null"`
    City         string `gorm:"not null"`
    Active       bool   `gorm:"not null"`
    RewardPoints float64    `gorm:"not null"`
}

func (*CustomerModel) Migrate(db *gorm.DB) error {
    return db.AutoMigrate(&CustomerModel{})
}
