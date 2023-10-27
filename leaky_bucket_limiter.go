package ratelimiter

import (
	"sync"
	"time"
)

type LeakyBucketLimiter struct {
	capacity int
	size     int
	duration time.Duration //
	lastTime time.Time     // the time of the last call to `Allow()`
	mu       *sync.Mutex
}

// AllowN capacity表示漏桶容量，超过容量则拒绝；每次请求时更新过去这段时间中剩下的水
func (l *LeakyBucketLimiter) AllowN(n int) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	duration := now.Sub(l.lastTime)

	// 计算自从上次请求后漏掉的水，减掉这部分水
	l.size = l.size - int(duration/l.duration)
	if l.size < 0 {
		l.size = 0
	}

	if l.size+n <= l.capacity {
		l.size += n
		return true
	}

	return false
}

func (l *LeakyBucketLimiter) Allow() bool {
	return l.AllowN(1)
}

func NewDefaultLeakyBucketLimiter() *LeakyBucketLimiter {
	return NewLeakyBucketLimiter(100, time.Millisecond*10)
}

func NewLeakyBucketLimiter(capacity int, duration time.Duration) *LeakyBucketLimiter {
	return &LeakyBucketLimiter{
		capacity: capacity,
		duration: duration,
		mu:       new(sync.Mutex),
	}
}
