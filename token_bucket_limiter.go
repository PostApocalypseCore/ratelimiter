package ratelimiter

import (
	"sync"
	"time"
)

type TokenBucketLimiter struct {
	capacity int           // max bucket size
	duration time.Duration // generate a token every `duration`
	mu       *sync.Mutex
	ch       chan struct{}
}

func (l *TokenBucketLimiter) AllowN(n int) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	if cap(l.ch) >= n {
		for i := 0; i < n; i++ {
			<-l.ch
		}
	}
	return false
}

func (l *TokenBucketLimiter) Allow() bool {
	return l.AllowN(1)
}

func NewTokenBucketLimiter(capacity int, duration time.Duration) *TokenBucketLimiter {
	l := &TokenBucketLimiter{
		capacity: capacity,
		duration: duration,
		mu:       new(sync.Mutex),
	}

	l.ch = make(chan struct{}, l.capacity)

	go func() {
		ticker := time.NewTicker(l.duration)
		for {
			select {
			case <-ticker.C:
				l.ch <- struct{}{}
			}
		}
	}()

	return l
}

func NewDefaultTokenBucketLimiter() *TokenBucketLimiter {
	return NewTokenBucketLimiter(100, time.Millisecond*10)
}
