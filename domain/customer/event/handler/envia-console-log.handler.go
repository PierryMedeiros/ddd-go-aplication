package handler

import (
    "fmt"
    eventShared "desafio-ddd-go/domain/1-shared/event"
)

type EnviaConsoleLogHandler struct {
    Address string
}

func (h *EnviaConsoleLogHandler) Handle(event eventShared.EventInterface) {
    
    eventData := event.EventData().(struct {
        ID      string
        Name    string
        Address struct {
            Street string
        }
    })

    fmt.Printf("Endereço do cliente: %s, %s alterado para: %s\n", eventData.ID, eventData.Name, eventData.Address.Street)
    h.Address = eventData.Address.Street
}
