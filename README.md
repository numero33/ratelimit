# numero33/ratelimit

Limit the rate of operations per time unit, use a [time.Ticker](https://golang.org/pkg/time/#NewTicker). This implementation refills the bucket based on the time elapsed between requests.

## Usage
Create new bucket
Refills every 18 seconds with 3 new requests and the bucket limit is 25

```shell
go get github.com/numero33/ratelimit
```

```golang
lr, err := ratelimit.NewLimiter(WithDuration(18*time.Second), WithAmount(3), WithLimit(25))
if err != nil {
    log.Fatal("ratelimit.NewLimiter")
}
```

to get a request

```golang
log.print("Wait for Slot")
lr.Take()
```

