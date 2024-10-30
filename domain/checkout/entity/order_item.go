package entity

type OrderItem struct {
    ID        string
    ProductID string
    Name      string
    Price     float64
    Quantity  int
    Total     float64
}

func NewOrderItem(id string, name string, price float64, productId string, quantity int) *OrderItem {
    return &OrderItem{
        ID:        id,
        Name:      name,
        Price:     price,
        ProductID: productId,
        Quantity:  quantity,
        Total:     price * float64(quantity),
    }
}

func (oi *OrderItem) GetID() string {
    return oi.ID
}

func (oi *OrderItem) GetName() string {
    return oi.Name
}

func (oi *OrderItem) GetProductID() string {
    return oi.ProductID
}

func (oi *OrderItem) GetQuantity() int {
    return oi.Quantity
}

func (oi *OrderItem) GetPrice() float64 {
    return oi.Price
}

func (oi *OrderItem) CalculateTotal() float64 {
    return oi.Price * float64(oi.Quantity)
}
