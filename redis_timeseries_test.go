package redis_timeseries

import (
	"testing"
	"time"
)

type columnTest struct {
	Interval time.Duration
	Times    []time.Time
	Results  []int64
}

func TestGet(t *testing.T) {
	testCases := []columnTest{
		columnTest{
			Interval: 1 * time.Minute,
			Times: []time.Time{
				time.Date(2000, time.January, 01, 00, 0, 30, 0, time.UTC),
			},
			Results: []int64{946684800},
		},

		columnTest{
			Interval: 15 * time.Minute,
			Times: []time.Time{
				time.Date(2000, time.January, 01, 00, 3, 0, 0, time.UTC),
				time.Date(2000, time.January, 01, 00, 16, 0, 0, time.UTC),
			},
			Results: []int64{946684800, 946685700},
		},

		columnTest{
			Interval: 23 * time.Minute,
			Times: []time.Time{
				time.Date(2000, time.January, 01, 00, 3, 0, 0, time.UTC),
				time.Date(2000, time.January, 01, 00, 23, 00, 0, time.UTC),
			},
			Results: []int64{946684800, 946686180},
		},

		columnTest{
			Interval: 15 * time.Minute,
			Times: []time.Time{
				time.Date(2000, time.January, 01, 00, 59, 59, 59, time.UTC),
			},
			Results: []int64{946687500},
		},
	}

	for _, column := range testCases {
		actualResults := Get(column.Interval, column.Times...)

		if len(actualResults) != len(column.Results) {
			t.Fatalf("expected %v results, got: %v",
				len(column.Results), len(actualResults),
			)
		}

		for index, expected := range column.Results {
			if expected != actualResults[index] {
				t.Fatalf("expected %v, got: %v, for: %v",
					expected, actualResults[index], column.Times[index],
				)
			}
		}
	}
}
