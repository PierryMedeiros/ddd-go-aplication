package entity

type ProductInterface interface {
    GetID() string
    GetName() string
    GetPrice() float64
}
