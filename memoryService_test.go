package simplemetrics_test

import (
	simplemetrics "github.com/co0p/simpleMetricsServiceInGo"
	"testing"
	"time"
)

func Test_InMemorySimpleMetricsService_Aggregate(t *testing.T) {

	e := simplemetrics.Event{"Label", 44, asTime("01.01.2020 14:04:00")}    // wrong time
	ea := simplemetrics.Event{"Another", 22, asTime("02.01.2020 15:00:00")} // wrong Label
	e1 := simplemetrics.Event{"Label", 2, asTime("01.01.2020 15:00:00")}
	e2 := simplemetrics.Event{"Label", 3, asTime("01.01.2020 15:00:05")}
	e3 := simplemetrics.Event{"Label", 5, asTime("01.01.2020 15:00:10")}
	e4 := simplemetrics.Event{"Label", 7, asTime("01.01.2020 15:00:20")}
	e5 := simplemetrics.Event{"Label", 13, asTime("01.01.2020 15:00:40")}
	e6 := simplemetrics.Event{"Label", 0, asTime("01.01.2020 15:00:50")}

	service := simplemetrics.InMemorySimpleMetricsService{}

	service.Record(e, ea, e1, e2, e3, e4, e5, e6)

	tenSecondsInMs := 10 * 1000

	buckets := service.Aggregate("Label", asTime("01.01.2020 15:00:00"), asTime("01.01.2020 15:01:00"), simplemetrics.AGGREGATE_SUM, tenSecondsInMs)

	actualLen := len(buckets)
	expectedLen := 7

	if actualLen != expectedLen {
		t.Errorf("expected len of bucket to be %d, got %d", expectedLen, actualLen)
	}

	counts := 0
	sum := 0
	for _, v := range buckets {
		counts += v.Count()
		sum += v.Sum()
	}

	expectedCounts := 6
	if counts != expectedCounts {
		t.Errorf("expected counts of events in buckets to be %d, got %d", expectedCounts, counts)
	}

	expectedSum := 30
	if sum != expectedSum {
		t.Errorf("expected sum of events in buckets to be %d, got %d", expectedSum, sum)
	}

}

func asTime(str string) time.Time {
	t, _ := time.Parse("02.01.2006 15:04:05", str)
	return t
}
