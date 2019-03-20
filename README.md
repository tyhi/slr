# tyhi/slr - Simple Rate Limiter

slr is a simple rate limiter I made for projects that don't require fancy rate limiting.

On windows this limit checker is zero cost as it never observes a clock tick. 

## Usage
```
callLimit := NewRateLimiter(100, time.Duration(time.Second*1), true)
for {
	start := time.Now()
	if callLimit.CheckLimit() == true {
		fmt.Println("Call limit hit.")
		return
	}
	log.Printf("Check took %s", time.Since(start))
	// do call
	
	fmt.Println(callLimit.Calls)
	time.Sleep(1 * time.Millisecond)
}
```