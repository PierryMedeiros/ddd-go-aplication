package factory

import (
    "desafio-ddd-go/domain/checkout/entity"
)

type OrderFactoryProps struct {
    ID         string
    CustomerID string
    Items      []OrderItemProps
}

type OrderItemProps struct {
    ID        string
    Name      string
    ProductID string
    Quantity  int
    Price     float64
}

type OrderFactory struct{}

func (OrderFactory) Create(props OrderFactoryProps) (*entity.Order, error) {
    items := []entity.OrderItem{}
    for _, item := range props.Items {
        orderItem := entity.OrderItem{
            ID:        item.ID,
            Name:      item.Name,
            ProductID: item.ProductID,
            Quantity:  item.Quantity,
            Price:     item.Price,
            Total:     item.Price * float64(item.Quantity),
        }
        items = append(items, orderItem)
    }

    order, err := entity.NewOrder(props.ID, props.CustomerID, items)
    if err != nil {
        return nil, err
    }
    return order, nil
}
