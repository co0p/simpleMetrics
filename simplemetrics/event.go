package simplemetrics

import (
	"fmt"
	"time"
)

type Event struct {
	Label        string
	Value        int // maybe int64
	OccurredDate time.Time
}

func NewEvent(label string, value int) Event {
	return Event{
		Label:        label,
		Value:        value,
		OccurredDate: time.Now(),
	}
}

func (e *Event) String() string {
	return fmt.Sprintf("Event: label='%s', value='%d' at %s", e.Label, e.Value, e.OccurredDate.Format("2006.01.02 15:04:05"))
}
