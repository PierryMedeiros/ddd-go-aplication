package event_test

import (
	eventDispatcher "desafio-ddd-go/domain/1-shared/event"
	cevent "desafio-ddd-go/domain/customer/event"
	chandler "desafio-ddd-go/domain/customer/event/handler"
	valueobject "desafio-ddd-go/domain/customer/value-object"
	"desafio-ddd-go/domain/product/event"
	"desafio-ddd-go/domain/product/event/handler"
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestEventDispatcher_Register(t *testing.T) {
	t.Log("Esta é uma mensagem de log")
	t.Logf("Valor de X: %d", 42)
	eventDispatcher := eventDispatcher.NewEventDispatcher()
	eventHandler := &handler.SendEmailWhenProductIsCreatedHandler{}

	eventDispatcher.Register("event.ProductCreatedEvent", eventHandler)

	handlers, ok := eventDispatcher.GetEventHandlers()["event.ProductCreatedEvent"]
	assert.True(t, ok)
	assert.Equal(t, 1, len(handlers))
	assert.Equal(t, eventHandler, handlers[0])
}

func TestEventDispatcher_Unregister(t *testing.T) {
	eventDispatcher := eventDispatcher.NewEventDispatcher()
	eventHandler := &handler.SendEmailWhenProductIsCreatedHandler{}

	eventDispatcher.Register("event.ProductCreatedEvent", eventHandler)
	handlers, ok := eventDispatcher.GetEventHandlers()["event.ProductCreatedEvent"]
	assert.True(t, ok)
	assert.Equal(t, eventHandler, handlers[0])

	eventDispatcher.Unregister("event.ProductCreatedEvent", eventHandler)
	handlers, ok = eventDispatcher.GetEventHandlers()["event.ProductCreatedEvent"]
	assert.True(t, ok)
	assert.Equal(t, 0, len(handlers))
}

func TestEventDispatcher_UnregisterAll(t *testing.T) {
	dispatcher := eventDispatcher.NewEventDispatcher()
	handler := &handler.SendEmailWhenProductIsCreatedHandler{}

	dispatcher.Register("productEvent.ProductCreatedEvent", handler)
	handlers, ok := dispatcher.GetEventHandlers()["productEvent.ProductCreatedEvent"]
	assert.True(t, ok)
	assert.Equal(t, handler, handlers[0])

	dispatcher.UnregisterAll()
	assert.Equal(t, 0, len(dispatcher.GetEventHandlers()))
}

func TestEventDispatcher_Notify(t *testing.T) {
	eventDispatcher := eventDispatcher.NewEventDispatcher()
	eventHandler := &handler.SendEmailWhenProductIsCreatedHandler{}
	
	eventDispatcher.Register("event.ProductCreatedEvent", eventHandler)
	handlers, ok := eventDispatcher.GetEventHandlers()["event.ProductCreatedEvent"]
	assert.True(t, ok)
	assert.Equal(t, eventHandler, handlers[0])

	productCreatedEvent := event.NewProductCreatedEvent(struct {
		Name        string
		Description string
		Price       float64
	}{
		Name:        "Product 1",
		Description: "Product 1 description",
		Price:       10.0,
	})

	eventDispatcher.Notify(productCreatedEvent)
}

func TestCustomerEvent(t *testing.T) {
	eventDispatcher := eventDispatcher.NewEventDispatcher()
	eventHandler := &chandler.EnviaConsoleLogHandler{}

	eventDispatcher.Register("event.CustomerCreatedEvent", eventHandler)
	handlers, ok := eventDispatcher.GetEventHandlers()["event.CustomerCreatedEvent"]
	assert.True(t, ok)
	assert.Equal(t, eventHandler, handlers[0])
	address, err := valueobject.NewAddress("Rua das Flores", 123, "12345-678", "São Paulo")
	if err != nil {
		fmt.Println("Erro ao criar endereço:", err)
		return
	}

	customerCreatedEvent := cevent.NewCustomerCreatedEvent(struct {
		ID      string
		Name    string
		Address interface{}
	}{
		ID:      "123",
		Name:    "Customer 1",
		Address: address,
	})

	eventDispatcher.Notify(customerCreatedEvent)
}
