package redis_timeseries

import (
	"time"

	"github.com/jinzhu/now"
)

func Get(interval time.Duration, times ...time.Time) []int64 {
	var results = []int64{}

	for _, t := range times {
		start := now.New(t).BeginningOfHour()

		for {
			// if t falls on start+interval, then we add start+interval
			if t.Equal(start.Add(interval)) {
				results = append(results, start.Add(interval).Unix())
				break
			}

			// if t isn't greater than start+interval, then we add start
			if !t.After(start.Add(interval)) {
				results = append(results, start.Unix())
				break
			}

			// this means t must be greater than start+interval
			start = start.Add(interval)
		}
	}

	return results
}
