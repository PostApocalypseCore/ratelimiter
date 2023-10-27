package ratelimiter

import (
	"sync"
	"time"
)

type FixedWindowLimiter struct {
	windowDuration  time.Duration // 一个窗口的时长大小
	maxRequestCount int           // 一个窗口允许多少请求
	windowId        int           // 当前窗口的id
	requestCount    int           // 当前窗口已有的请求个数
	mu              *sync.Mutex
}

func (l *FixedWindowLimiter) Allow() bool {
	id := int(time.Now().UnixNano() / int64(l.windowDuration))

	l.mu.Lock()
	defer l.mu.Unlock()

	if id > l.windowId {
		l.windowId = id
		l.requestCount = 0
	}
	if l.requestCount < l.maxRequestCount {
		l.requestCount += 1
		return true
	}
	return false
}

func NewFixedWindowLimiter(windowDuration time.Duration, maxRequestCount int) *FixedWindowLimiter {
	return &FixedWindowLimiter{
		windowDuration:  windowDuration,
		maxRequestCount: maxRequestCount,
		mu:              &sync.Mutex{},
		windowId:        0,
		requestCount:    0,
	}
}

func NewDefaultFixedWindowLimiter() *FixedWindowLimiter {
	return &FixedWindowLimiter{
		windowDuration:  time.Second,
		maxRequestCount: 100,
		mu:              &sync.Mutex{},
		windowId:        0,
		requestCount:    0,
	}
}
