package redis_timeseries

import "time"

// Get returns unix timestamps of the times closest to each time argument
// as specified by the interval.
//
// Ex:
//   newyear := time.Date(2000, time.January, 01, 0, 0, 30, 0, time.UTC)
//   newyearplus := time.Date(2000, time.January, 01, 0, 16, 0, 0, time.UTC)
//
//   // returns: []string{946684800, 946685700}
//   results := Get(15*time.Minute, newyear, newyearplus)
//
// `newyear` falls between 00:00 and 00:15, so it returns 00:00.
// `newyearplus` falls between 00:15 and 00:30, so it returns 00:15.
//
// If a time falls on the interval, ie 00:15, it returns 00.15.
func Get(interval time.Duration, times ...time.Time) []int64 {
	var results = []int64{}

	for _, t := range times {
		start := time.Date(2000, t.Month(), t.Day(), t.Hour(), 00, 0, 0, time.UTC)

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

			// this means t must be greater than start+interval, however
			// we don't know by how much, so we continue with the loop
			start = start.Add(interval)
		}
	}

	return results
}
