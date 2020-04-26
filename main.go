package main

import (
	"fmt"
	"time"
)

type Event struct {
	label       string
	value       int // maybe int64
	occuredDate time.Time
}

type Bucket struct {
	startDate time.Time
	events    []Event
}

func (b *Bucket) count() int {
	return len(b.events)
}

func (b *Bucket) sum() int {
	sum := 0

	for _, v := range b.events {
		sum += v.value
	}
	return sum
}

type Aggregation string

const (
	AGGREGATE_SUM  Aggregation = "sum"
	AGGREGATE_MEAN Aggregation = "mean"
)

type SimpleMetricsService interface {
	Record(...Event)
	Aggregate(label string, startDate time.Time, endDate time.Time, aggregation Aggregation, sampleRate int) []Bucket
}

type InMemorySimpleMetricsService struct {
	events []Event
}

func (m *InMemorySimpleMetricsService) Record(events ...Event) {
	for _, v := range events {
		m.events = append(m.events, v)
	}
}

func (m *InMemorySimpleMetricsService) Aggregate(label string, startDate time.Time, endDate time.Time, aggregation Aggregation, sampleRateMS int) []Bucket {

	events := m.getRelevantEvents(label, startDate, endDate)

	fmt.Printf("we have %d events!\n", len(events))

	buckets := []Bucket{}
	bucketStartDate := startDate

	for bucketStartDate.Before(endDate) || bucketStartDate.Equal(endDate) {
		nextBucketDate := bucketStartDate.Add(time.Duration(sampleRateMS) * time.Millisecond)

		bucket := Bucket{
			startDate: bucketStartDate,
			events:    []Event{},
		}
		fmt.Println("current bucket", bucket)

		for k, _ := range events {
			event := events[k]

			// in range?
			if event.occuredDate.Equal(bucket.startDate) ||
				event.occuredDate.After(bucket.startDate) &&
					event.occuredDate.Before(nextBucketDate) {

				fmt.Println("adding event to bucket", event)
				bucket.events = append(bucket.events, event)
			}
		}

		buckets = append(buckets, bucket)
		bucketStartDate = nextBucketDate
	}
	// ...

	// 3. aggregate events to corresponding bucket
	// ...
	return buckets
}

func (m *InMemorySimpleMetricsService) getRelevantEvents(label string, startDate time.Time, endDate time.Time) []Event {
	events := []Event{}

	for _, candidate := range m.events {

		if candidate.label == label &&
			(candidate.occuredDate.Equal(startDate) || candidate.occuredDate.After(startDate)) &&
			candidate.occuredDate.Before(endDate) {

			events = append(events, candidate)
		}
	}

	return events
}

func main() {

	// events:
	// 00 			10 	20 	30	40	50	60
	// sum(e1,e1)	e3	e4	__	e5	e6  __
	// => [5, 5, 7, 0, 13, 0, 0]
	// count:
	// => [2, 1, 1, 0,  1, 1, 0]

	e := Event{"label", 44, asTime("01.01.2020 14:04:00")}    // wrong time
	ea := Event{"Another", 22, asTime("02.01.2020 15:00:00")} // wrong label
	e1 := Event{"label", 2, asTime("01.01.2020 15:00:00")}
	e2 := Event{"label", 3, asTime("01.01.2020 15:00:05")}
	e3 := Event{"label", 5, asTime("01.01.2020 15:00:10")}
	e4 := Event{"label", 7, asTime("01.01.2020 15:00:20")}
	e5 := Event{"label", 13, asTime("01.01.2020 15:00:40")}
	e6 := Event{"label", 0, asTime("01.01.2020 15:00:50")}

	service := InMemorySimpleMetricsService{}

	service.Record(e, ea, e1, e2, e3, e4, e5, e6)

	tenSecondsInMs := 10 * 1000

	buckets := service.Aggregate("label", asTime("01.01.2020 15:00:00"), asTime("01.01.2020 15:01:00"), AGGREGATE_SUM, tenSecondsInMs)

	for k, v := range buckets {
		fmt.Printf("%d: \t value: %v, count: %d\n", k, v.sum(), v.count)
	}
}

func asTime(str string) time.Time {
	t, _ := time.Parse("02.01.2006 15:04:05", str)
	return t
}
