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

func (m *InMemorySimpleMetricsService) Aggregate(query Query) []Bucket {

	events := m.getRelevantEvents(query.label, query.startDate, query.endDate)

	fmt.Printf("we have %d events!\n", len(events))

	buckets := []Bucket{}
	bucketStartDate := query.startDate

	for bucketStartDate.Before(query.endDate) || bucketStartDate.Equal(query.endDate) {
		nextBucketDate := bucketStartDate.Add(time.Duration(query.sampleRate) * time.Millisecond)

		bucket := Bucket{
			startDate: bucketStartDate,
			events:    []Event{},
		}
		for k, _ := range events {
			event := events[k]

			// in range?
			if event.OccurredDate.Equal(bucket.startDate) ||
				event.OccurredDate.After(bucket.startDate) &&
					event.OccurredDate.Before(nextBucketDate) {

				fmt.Println("adding event to bucket", event)
				bucket.events = append(bucket.events, event)
			}
		}

		buckets = append(buckets, bucket)
		bucketStartDate = nextBucketDate
	}

	return buckets
}

func (m *InMemorySimpleMetricsService) getRelevantEvents(label string, startDate time.Time, endDate time.Time) []Event {
	events := []Event{}

	for _, candidate := range m.events {

		if candidate.Label == label &&
			(candidate.OccurredDate.Equal(startDate) || candidate.OccurredDate.After(startDate)) &&
			candidate.OccurredDate.Before(endDate) {

			events = append(events, candidate)
		}
	}

	return events
}
