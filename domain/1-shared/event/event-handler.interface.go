package event

type EventHandlerInterface interface {
    Handle(event EventInterface)
}