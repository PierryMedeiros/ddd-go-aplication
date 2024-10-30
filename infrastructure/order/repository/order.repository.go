package repository

import (
	"errors"
	"gorm.io/gorm"
	"desafio-ddd-go/domain/checkout/entity"
	"desafio-ddd-go/infrastructure/models"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(order *entity.Order) error {
	orderModel := models.OrderModel{
		ID:         order.GetID(),
		CustomerID: order.GetCustomerID(),
		Total:      order.Total,
		Items:      make([]models.OrderItemModel, len(order.GetItems())),
	}

	for i, item := range order.GetItems() {
		orderModel.Items[i] = models.OrderItemModel{
			ID:        item.GetID(),
			Name:      item.GetName(),
			Price:     item.GetPrice(),
			ProductID: item.GetProductID(),
			Quantity:  item.GetQuantity(),
		}
	}

	return r.db.Create(&orderModel).Error
}

func (r *OrderRepository) Update(order *entity.Order) error {
	if err := r.db.Model(&models.OrderModel{}).Where("id = ?", order.GetID()).Updates(models.OrderModel{
		CustomerID: order.GetCustomerID(),
		Total:      order.Total,
	}).Error; err != nil {
		return err
	}

	if err := r.db.Where("order_id = ?", order.GetID()).Delete(&models.OrderItemModel{}).Error; err != nil {
		return err
	}

	for _, item := range order.GetItems() {
		orderItemModel := models.OrderItemModel{
			ID:        item.GetID(),
			Name:      item.GetName(),
			Price:     item.GetPrice(),
			ProductID: item.GetProductID(),
			Quantity:  item.GetQuantity(),
			OrderID:   order.GetID(),
		}
		if err := r.db.Create(&orderItemModel).Error; err != nil {
			return err
		}
	}

	return nil
}

func (r *OrderRepository) Find(id string) (*entity.Order, error) {
	var orderModel models.OrderModel
	if err := r.db.Preload("Items").First(&orderModel, "id = ?", id).Error; err != nil {
		return nil, errors.New("not found")
	}

	items := make([]entity.OrderItem, len(orderModel.Items))
	for i, item := range orderModel.Items {
		items[i] = *entity.NewOrderItem(item.ID, item.Name, item.Price, item.ProductID, item.Quantity)
	}

	order, err := entity.NewOrder(orderModel.ID, orderModel.CustomerID, items)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *OrderRepository) FindAll() ([]*entity.Order, error) {
	var orderModels []models.OrderModel
	if err := r.db.Preload("Items").Find(&orderModels).Error; err != nil {
		return nil, err
	}

	var orders []*entity.Order
	for _, orderModel := range orderModels {
		items := make([]entity.OrderItem, len(orderModel.Items))
		for i, item := range orderModel.Items {
			items[i] = *entity.NewOrderItem(item.ID, item.Name, item.Price, item.ProductID, item.Quantity)
		}

		order, err := entity.NewOrder(orderModel.ID, orderModel.CustomerID, items)
		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}
