package slr

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestNewRateLimiter(t *testing.T) {
	callLimit := NewRateLimiter(100, time.Duration(time.Second*1), true)
	for {
		start := time.Now()
		if callLimit.CheckLimit() == true {
			fmt.Println("Call limit hit.")
			return
		}
		log.Printf("Check took %s", time.Since(start))

		fmt.Println(callLimit.Calls)
		time.Sleep(1 * time.Millisecond)
	}
}
