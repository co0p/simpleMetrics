package simplemetrics

import (
	"fmt"
	"time"
)

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
			if event.OccuredDate.Equal(bucket.startDate) ||
				event.OccuredDate.After(bucket.startDate) &&
					event.OccuredDate.Before(nextBucketDate) {

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

		if candidate.Label == label &&
			(candidate.OccuredDate.Equal(startDate) || candidate.OccuredDate.After(startDate)) &&
			candidate.OccuredDate.Before(endDate) {

			events = append(events, candidate)
		}
	}

	return events
}
