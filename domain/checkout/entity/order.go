package entity

import "errors"

type Order struct {
    ID         string
    CustomerID string
    Items      []OrderItem
    Total      float64
}

func NewOrder(id string, customerId string, items []OrderItem) (*Order, error) {
    o := &Order{
        ID:         id,
        CustomerID: customerId,
        Items:      items,
    }
    o.Total = o.calculateTotal()
    if err := o.Validate(); err != nil {
        return nil, err
    }
    return o, nil
}

func (o *Order) GetID() string {
    return o.ID
}

func (o *Order) GetCustomerID() string {
    return o.CustomerID
}

func (o *Order) GetItems() []OrderItem {
    return o.Items
}

func (o *Order) Validate() error {
    if len(o.ID) == 0 {
        return errors.New("id is required")
    }
    if len(o.CustomerID) == 0 {
        return errors.New("customerId is required")
    }
    if len(o.Items) == 0 {
        return errors.New("items are required")
    }
    for _, item := range o.Items {
        if item.Quantity <= 0 {
            return errors.New("quantity must be greater than 0")
        }
    }
    return nil
}

func (o *Order) calculateTotal() float64 {
    var total float64
    for _, item := range o.Items {
        total += item.Total
    }
    return total
}
