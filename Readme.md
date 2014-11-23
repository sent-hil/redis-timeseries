# redis_timeseries

redis_timeseries is an library that helps storing timeseries in redis. This library doesn't actually talk to redis, but returns the keys that you can then use with a redis client like [redigo](https://github.com/garyburd/redigo) or [go-redis](https://github.com/go-redis/redis).

## Install

```
go get -u github.com/sent-hil/redis_timeseries
```

## Usage

redis_timeseries assumes use of sets or sorted sets to store the values. For example, you want to build a timeseries with interval of 15 minutes, these are the guidelines:

* timestamps from t=0 till t=15 belongs to t=0 bucket
* timestamps from from t=15 till t=30 belongs to t=15 bucket
* timestamps that occurs t=15 belongs to t=15 bucket

`redis_timeseries#Get()` accepts `time.Duration` as first argument and a variadic number of `time.Time` arguments and returns an array of bucket timestamps. Each result corresponding to the index location of argument.

```go
// when: time.Now() -> 2000-01-01 00:00:30 +0000 UTC
// returns: []string{946684800}
redis_timeseries.Get(1*time.Minute, time.Now())

// when: time.Now() -> 2000-01-01 00:03:00 +0000
// returns: []string{946684800, 946685700}
redis_timeseries.Get(15*time.Minute, time.Now(), time.Now().Add(16*time.Minute))
```
