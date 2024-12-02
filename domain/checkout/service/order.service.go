package service

import (
    "errors"
    "github.com/google/uuid"
    "desafio-ddd-go/domain/checkout/entity"
    customer "desafio-ddd-go/domain/customer/entity"
)

type OrderService struct{}

func (s OrderService) PlaceOrder(customer *customer.Customer, items []entity.OrderItem) (*entity.Order, error) {
    if len(items) == 0 {
        return nil, errors.New("order must have at least one item")
    }

    orderID := uuid.New().String()
    order, err := entity.NewOrder(orderID, customer.GetId(), items)
    if err != nil {
        return nil, err
    }

    customer.AddRewardPoints(order.GetTotal() / 2)
    return order, nil
}

func (s OrderService) Total(orders []entity.Order) float64 {
    var total float64
    for _, order := range orders {
        total += order.GetTotal()
    }
    return total
}
