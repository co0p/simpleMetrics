package simplemetrics

import "time"

type SimpleMetricsService interface {
	Record(...Event)
	Aggregate(label string, startDate time.Time, endDate time.Time, aggregation Aggregation, sampleRate int) []Bucket
}
