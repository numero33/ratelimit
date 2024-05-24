package ratelimit

import (
	"time"
)

type Limiter struct {
	duration time.Duration
	amount   uint64
	limit    uint64
	slot     chan bool
	ticker   *time.Ticker
	preFill  uint64
}

func NewLimiter(options ...LimiterOptionFunc) (*Limiter, error) {

	l := &Limiter{}

	for _, fn := range options {
		if fn == nil {
			continue
		}
		if err := fn(l); err != nil {
			return nil, err
		}
	}

	l.slot = make(chan bool, l.limit)

	if l.preFill != 0 {
		if l.preFill > l.limit {
			l.preFill = l.limit
		}
		for i := uint64(0); i < l.preFill; i++ {
			l.slot <- true
		}
	}

	if l.duration != 0 {
		l.ticker = time.NewTicker(l.duration)
		go func(b *Limiter) {
			for range b.ticker.C {
				b.refill()
			}
		}(l)
	}

	return l, nil
}

func (l *Limiter) Take() {
	<-l.slot
}

func (l *Limiter) refill() {

	amount := l.amount
	if l.limit != 0 && uint64(len(l.slot))+amount > l.limit {
		amount = l.limit - uint64(len(l.slot))
	}

	for i := uint64(0); i < amount; i++ {
		l.slot <- true
	}
}

// LimiterOptionFunc can be used to customize a new Bucket
type LimiterOptionFunc func(limiter *Limiter) error

// WithPreFill fills the bucket with the specified amount of slots
func WithPreFill(count uint64) LimiterOptionFunc {
	return func(l *Limiter) error {
		l.preFill = count
		return nil
	}
}

// WithLimit sets the limit for the bucket
func WithLimit(count uint64) LimiterOptionFunc {
	return func(l *Limiter) error {
		l.limit = count
		return nil
	}
}

// WithAmount sets the amount of slots to refill
func WithAmount(amount uint64) LimiterOptionFunc {
	return func(l *Limiter) error {
		l.amount = amount
		return nil
	}
}

// WithDuration sets the duration for the refill
func WithDuration(duration time.Duration) LimiterOptionFunc {
	return func(l *Limiter) error {
		l.duration = duration
		return nil
	}
}
