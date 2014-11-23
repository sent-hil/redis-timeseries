# redis-timeseries

redis-timeseries is an library that helps storing timeseries in redis. This library doesn't actually talk to redis, but returns the keys that you can then use with a redis client like [redigo](https://github.com/garyburd/redigo) or [go-redis](https://github.com/go-redis/redis).

## Install

```
go get -u github.com/sent-hil/redis-timeseries
```

redis-timeseries uses redis' sorted sets to store the values. For example, you want to build a timeseries with interval of 15 minutes. Any data that comes from t=0 till t=15 will be stored and incremented in t=0 bucket.

```go
```

You can specify a string to be prepended to each key like:

```go
```

To retreive a combined results of the sets (using redis' `ZINTERSTORE` command) involved:

```go
```
