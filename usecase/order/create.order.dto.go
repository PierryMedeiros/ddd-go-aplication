package order

type CreateOrderDto struct {
	OrderID    string                   
	CustomerID string                   
	Items      []CreateOrderItemDto 
}

type CreateOrderItemDto struct {
	ID        string  
	Name      string  
	Price     float64 
	ProductID string  
	Quantity  int     
}