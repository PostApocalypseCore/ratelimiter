package ratelimiter

import "testing"

func TestTokenBucketLimiter_Allow(t *testing.T) {
	test1(NewDefaultTokenBucketLimiter(), t)
}
