package simplemetrics

type SimpleMetricsService interface {
	Record(...Event)
	Aggregate(Query) []Bucket
}
