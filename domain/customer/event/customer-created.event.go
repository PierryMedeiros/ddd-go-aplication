package event

import "time"

type CustomerCreatedEvent struct {
    dataTimeOccurred time.Time
    eventData        interface{}
}

func NewCustomerCreatedEvent(eventData interface{}) *CustomerCreatedEvent {
    return &CustomerCreatedEvent{
        dataTimeOccurred: time.Now(),
        eventData:        eventData,
    }
}

func (e *CustomerCreatedEvent) DataTimeOccurred() time.Time {
    return e.dataTimeOccurred
}

func (e *CustomerCreatedEvent) EventData() interface{} {
    return e.eventData
}
