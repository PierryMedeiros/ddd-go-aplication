package event

type EventDispatcherInterface interface {
  Register(eventName string, eventHandler EventHandlerInterface)
  Unregister(eventName string, eventHandler EventHandlerInterface)
  UnregisterAll()
  Notify(event EventInterface)
}