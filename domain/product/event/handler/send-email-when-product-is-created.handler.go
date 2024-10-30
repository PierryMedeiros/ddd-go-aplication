package handler

import (
    "fmt"
    eventShared "desafio-ddd-go/domain/1-shared/event"
)
 
type SendEmailWhenProductIsCreatedHandler struct{}

func (h *SendEmailWhenProductIsCreatedHandler) Handle(event eventShared.EventInterface) {
    fmt.Println("Sending email to .....")
}
