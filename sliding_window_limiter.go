package ratelimiter

import (
	"container/list"
	"sync"
	"time"
)

type SlideWindowLimiter struct {
	windowDuration  time.Duration // 一个窗口的时长大小
	maxRequestCount int           // 一个窗口允许多少请求
	mu              *sync.Mutex
	list            *list.List // 存储时间戳
}

func NewDefaultSlideWindowLimiter() *SlideWindowLimiter {
	return &SlideWindowLimiter{
		windowDuration:  time.Second,
		maxRequestCount: 100,
		mu:              &sync.Mutex{},
		list:            list.New(),
	}
}

// Allow 检查当前窗口内的请求数（从当前时刻回溯一个窗口），这要求我们记录之前的请求的时间戳
func (l *SlideWindowLimiter) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	for i := l.list.Front(); i != nil; i = i.Next() {
		reqTime := i.Value.(time.Time)
		if now.Sub(reqTime) < l.windowDuration {
			break // 遍历到当前窗口后退出
		}
		l.list.Remove(i) //将之前的窗口的请求的时间戳全部清除
	}
	if l.list.Len()+1 <= l.maxRequestCount {
		l.list.PushBack(now)
		//log.Println("allow: ", now.UnixMilli())
		return true
	}
	//log.Println("deny: ", now.UnixMilli())
	return false
}
