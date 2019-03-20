package slr

import (
	"time"
)

// rateLimiter holds the information about the rate limiter that is created.
type rateLimiter struct {
	CurrentStamp time.Time
	Interval     time.Duration
	Wait         bool
	AllowedCalls int
	Calls        int
}

// NewRateLimiter creates a new rate limiter while added the current time to the rateLimit struct.
func NewRateLimiter(calls int, duration time.Duration, wait bool) *rateLimiter {
	return &rateLimiter{CurrentStamp: time.Now(), Interval: duration, Calls: 0, AllowedCalls: calls, Wait: wait}
}

// CheckLimit will check to see if we've hit the rate limit and either return false or wait until the interval passes.
func (rl *rateLimiter) CheckLimit() bool {
	if time.Since(rl.CurrentStamp) >= rl.Interval {
		resetInterval(rl)
		rl.Calls++
		return false
	}
	if rl.Calls >= rl.AllowedCalls {
		if rl.Wait == true {
			time.Sleep(rl.Interval - time.Since(rl.CurrentStamp))
			resetInterval(rl)
			rl.Calls++
			return false
		}
		return true
	}
	rl.Calls++
	return false
}

// resetInterval will change the currentstamp to Now() and reset the number of calls.
func resetInterval(rl *rateLimiter) {
	rl.CurrentStamp = time.Now()
	rl.Calls = 0
}
