package event

import "time"

type EventInterface interface {
    DataTimeOccurred() time.Time
    EventData() interface{}
}   