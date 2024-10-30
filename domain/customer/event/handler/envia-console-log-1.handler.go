package handler

import (
    "fmt"
    "desafio-ddd-go/domain/customer/event"
)

type EventHandlerInterface interface {
    Handle(event CustomerCreatedEvent)
}

type CustomerCreatedEvent struct {}

type EnviaConsoleLog1Handler struct{}

func (h *EnviaConsoleLog1Handler) Handle(event event.CustomerCreatedEvent) {
    fmt.Println("Esse é o primeiro console.log do evento: CustomerCreated")
}
