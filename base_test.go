package ratelimiter

import (
	"testing"
	"time"
)

func test1(l IRateLimiter, t *testing.T) {
	t.Helper()
	ticker := time.NewTicker(time.Millisecond * 10)
	forever := true
	for i := 0; forever; {
		select {
		case <-ticker.C:
			i++
			if i <= 100 && !l.Allow() {
				t.Fatalf("limiter err")
			}
			if i > 100 {
				forever = false
			}
		}
	}
}

func test2(l IRateLimiter, t *testing.T) {
	t.Helper()
	ticker := time.NewTicker(time.Millisecond * 5)
	for i, forever := 0, true; forever; {
		select {
		case <-ticker.C:
			i++
			ok := l.Allow()
			if i <= 100 && !ok {
				t.Fatalf("limiter err 1")
			}
			if i > 100 && ok {
				t.Fatalf("limiter err 2")
			}
			if i >= 200 {
				forever = false
			}
		}
	}
}
