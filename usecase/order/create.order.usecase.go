package order

import (
	orderEntity "desafio-ddd-go/domain/checkout/entity"
	"desafio-ddd-go/infrastructure/models"
	"desafio-ddd-go/infrastructure/order/repository"
	"gorm.io/gorm"
	"log"
)

type CreateOrderUseCase struct {
	OrderID    string
	CustomerID string
	Items      []CreateOrderItemDto
	db         *gorm.DB
}

func NewCreateOrderUseCase(props CreateOrderDto, db *gorm.DB) *CreateOrderUseCase {

	return &CreateOrderUseCase{
		OrderID:    props.OrderID,
		CustomerID: props.CustomerID,
		Items:      props.Items,
		db:         db,
	}
}

func (c *CreateOrderUseCase) Execute() error {

	repo := repository.NewOrderRepository(c.db)

	var orderItems []orderEntity.OrderItem

	for _, item := range c.Items {
		var product models.ProductModel
		if err := c.db.First(&product, "id = ?", item.ProductID).Error; err != nil {
			log.Printf("Produto não encontrado: %v", err)
			return err
		}

		orderItem := orderEntity.NewOrderItem(item.ID, item.Name, item.Price, item.ProductID, item.Quantity)
		orderItems = append(orderItems, *orderItem)
	}

	order, err := orderEntity.NewOrder(c.OrderID, c.CustomerID, orderItems)

	if err != nil {
		log.Printf("Erro ao ceiar ordem: %v", err)
	}
	log.Printf("Erro ao ceiar ordem: %v", err)
	if err := repo.Create(order); err != nil {
		log.Printf("Erro ao salvar pedido no banco: %v", err)
		return err
	}

	return nil
}
