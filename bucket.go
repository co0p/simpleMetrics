package simplemetrics

import "time"

type Bucket struct {
	startDate time.Time
	events    []Event
}

func (b *Bucket) Count() int {
	return len(b.events)
}

func (b *Bucket) Sum() int {
	sum := 0

	for _, v := range b.events {
		sum += v.Value
	}
	return sum
}
