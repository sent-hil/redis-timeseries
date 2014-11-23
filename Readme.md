# redis_timeseries

redis_timeseries is an library that helps storing timeseries in redis. This library doesn't actually talk to redis, but returns the keys that you can then use with a redis client like [redigo](https://github.com/garyburd/redigo) or [go-redis](https://github.com/go-redis/redis).

## Install

```
go get -u github.com/sent-hil/redis_timeseries
```

redis_timeseries uses redis' sorted sets to store the values. For example, you want to build a timeseries with interval of 15 minutes. Any data that comes from t=0 till t=15 will be stored and incremented in t=0 bucket.

```go
// when: time.Now() -> Sun, 23 Nov 2014 00:00:30 GMT
// returns: []string{1416700800}
redis_timeseries.Get(1*time.Minute, time.Now())

// when: time.Now() -> Sun, 23 Nov 2014 00:01:00 GMT
// returns: []string{1416700800, 1416701700}
redis_timeseries.Get(15*time.Minute, time.Now(), time.Now().Add(16*time.Minute))
```
