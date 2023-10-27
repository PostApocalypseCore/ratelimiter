package ratelimiter

type IRateLimiter interface {
	Allow() bool
}
