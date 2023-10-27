package ratelimiter

import (
	"testing"
)

func TestFixedWindowLimiter_Limit(t *testing.T) {
	test1(NewDefaultFixedWindowLimiter(), t)
}
