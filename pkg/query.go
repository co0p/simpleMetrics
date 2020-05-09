package simplemetrics

import "time"

type Query struct {
	label       string
	startDate   time.Time
	endDate     time.Time
	aggregation Aggregation
	sampleRate  int
}

func NewQuery(label string, startDate time.Time, endDate time.Time, aggregation Aggregation, sampleRate int) Query {

	// todo some validation

	return Query{
		label:       label,
		startDate:   startDate,
		endDate:     endDate,
		aggregation: aggregation,
		sampleRate:  sampleRate,
	}
}
