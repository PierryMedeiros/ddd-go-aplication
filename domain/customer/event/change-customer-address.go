package event

import "time"

type CustomerChangeAddressEvent struct {
    dataTimeOccurred time.Time
    eventData        struct {
        ID      string
        Name    string
        Address struct {
            Street string
        }
    }
}

func NewCustomerChangeAddressEvent(eventData struct {
    ID      string
    Name    string
    Address struct {
        Street string
    }

}) *CustomerChangeAddressEvent {
    return &CustomerChangeAddressEvent{
        dataTimeOccurred: time.Now(),
        eventData:        eventData,
    }
}

func (e *CustomerChangeAddressEvent) DataTimeOccurred() time.Time {
    return e.dataTimeOccurred
}

func (e *CustomerChangeAddressEvent) EventData() interface{} {
    return e.eventData
}