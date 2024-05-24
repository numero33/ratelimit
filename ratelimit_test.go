package ratelimit

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewBucket(t *testing.T) {

	duration := time.Second * 1
	refill := uint64(5)
	limit := uint64(20)
	prefill := uint64(10)

	l, err := NewLimiter(WithDuration(duration), WithAmount(refill), WithLimit(limit), WithPreFill(prefill))

	assert.Equal(t, l.duration, duration)
	assert.Equal(t, l.limit, limit)
	assert.Equal(t, l.amount, refill)
	assert.Equal(t, uint64(len(l.slot)), prefill)
	assert.Nil(t, err)
}

func TestNewBucket_Speed(t *testing.T) {
	rl, _ := NewLimiter(WithLimit(100), WithPreFill(10))

	last := time.Now()
	for i := 0; i < 10; i++ {
		rl.Take()
		cur := time.Now()
		fmt.Println("last", cur.Sub(last))
		last = cur
	}
}
