package simplemetrics

import "time"

type Bucket struct {
	startDate time.Time
	events    []Event
}

func (b *Bucket) Count() int {
	return len(b.events)
}

// TODO: move to aggregation instead, bucket should know nothing about applied aggregation on itself
func (b *Bucket) Sum() int {
	sum := 0

	for _, v := range b.events {
		sum += v.Value
	}
	return sum
}
