package handler

import (
    "fmt"
    "desafio-ddd-go/domain/customer/event"
)

type EventHandlerInterface2 interface {
    Handle(event CustomerCreatedEvent)
}

type CustomerCreatedEvent2 struct {}

type EnviaConsoleLog1Handler2 struct{}

func (h *EnviaConsoleLog1Handler) Handle2(event event.CustomerCreatedEvent) {
    fmt.Println("Esse é o segundo console.log do evento: CustomerCreated")
}
