package ratelimiter

import (
	"testing"
)

func TestSlideWindowLimiter_Allow(t *testing.T) {
	test1(NewDefaultSlideWindowLimiter(), t)
}

func TestSlideWindowLimiter_Allow2(t *testing.T) {
	test2(NewDefaultSlideWindowLimiter(), t)
}
