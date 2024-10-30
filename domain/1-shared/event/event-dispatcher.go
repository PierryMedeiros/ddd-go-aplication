package event

import "fmt"

type EventDispatcher struct {
    eventHandlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
    return &EventDispatcher{
        eventHandlers: make(map[string][]EventHandlerInterface),
    }
}

func (ed *EventDispatcher) GetEventHandlers() map[string][]EventHandlerInterface {
    return ed.eventHandlers
}

func (ed *EventDispatcher) Register(eventName string, eventHandler EventHandlerInterface) {
    if _, ok := ed.eventHandlers[eventName]; !ok {
        ed.eventHandlers[eventName] = []EventHandlerInterface{}
    }
    ed.eventHandlers[eventName] = append(ed.eventHandlers[eventName], eventHandler)
}

func (ed *EventDispatcher) Unregister(eventName string, eventHandler EventHandlerInterface) {
    if handlers, ok := ed.eventHandlers[eventName]; ok {
        for i, handler := range handlers {
            if handler == eventHandler {
                ed.eventHandlers[eventName] = append(handlers[:i], handlers[i+1:]...)
                break
            }
        }
    }
}

func (ed *EventDispatcher) UnregisterAll() {
    ed.eventHandlers = make(map[string][]EventHandlerInterface)
}

func (ed *EventDispatcher) Notify(event EventInterface) {
    eventName := fmt.Sprintf("%T", event)
    if handlers, ok := ed.eventHandlers[eventName]; ok {
        for _, handler := range handlers {
            handler.Handle(event)
        }
    }
}
