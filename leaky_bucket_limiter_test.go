package ratelimiter

import (
	"testing"
)

func TestLeakyBucketLimiter_Allow(t *testing.T) {
	test1(NewDefaultLeakyBucketLimiter(), t)
}
