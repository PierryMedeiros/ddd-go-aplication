package service

import (
	"desafio-ddd-go/domain/checkout/entity"
	customer "desafio-ddd-go/domain/customer/entity"
	"github.com/google/uuid"
	"testing"
)

func TestOrderService_PlaceOrder(t *testing.T) {
	customer, err := customer.NewCustomer("1", "Customer 1")
  if err != nil { t.Error(err) }

	items := []entity.OrderItem{
		{
			ID:        uuid.New().String(),
			ProductID: uuid.New().String(),
			Quantity:  1,
			Price:     100.0,
			Total:     100.0,
		},
	}

	orderService := OrderService{}

	order, err := orderService.PlaceOrder(customer, items)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if order == nil {
		t.Fatal("expected order, got nil")
	}

	if order.GetTotal() != 100.0 {
		t.Errorf("expected total 100.0, got %v", order.GetTotal())
	}

	expectedRewardPoints := 50.0
	if customer.GetRewardPoints() != expectedRewardPoints {
		t.Errorf("expected reward points %v, got %v", expectedRewardPoints, customer.GetRewardPoints())
	}
}

func TestOrderService_Total(t *testing.T) {
	orders := []entity.Order{
		{
			ID:         uuid.New().String(),
			CustomerID: uuid.New().String(),
			Items: []entity.OrderItem{
				{
					ID:        uuid.New().String(),
					ProductID: uuid.New().String(),
					Quantity:  1,
					Price:     100.0,
					Total:     100.0,
				},
			},
			Total: 100.0,
		},
		{
			ID:         uuid.New().String(),
			CustomerID: uuid.New().String(),
			Items: []entity.OrderItem{
				{
					ID:        uuid.New().String(),
					ProductID: uuid.New().String(),
					Quantity:  2,
					Price:     50.0,
					Total:     100.0,
				},
			},
			Total: 100.0,
		},
	}

	orderService := OrderService{}
	total := orderService.Total(orders)

	expectedTotal := 200.0
	if total != expectedTotal {
		t.Errorf("expected total %v, got %v", expectedTotal, total)
	}
}
