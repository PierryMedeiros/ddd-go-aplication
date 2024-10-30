package event

import "time"

type ProductCreatedEvent struct {
    dataTimeOccurred time.Time
    eventData        interface{}
}

func NewProductCreatedEvent(eventData interface{}) *ProductCreatedEvent {
    return &ProductCreatedEvent{
        dataTimeOccurred: time.Now(),
        eventData:        eventData,
    }
}

func (e *ProductCreatedEvent) DataTimeOccurred() time.Time {
    return e.dataTimeOccurred
}

func (e *ProductCreatedEvent) EventData() interface{} {
    return e.eventData
}
