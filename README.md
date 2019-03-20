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

## Todo
* Avoid a wait race condition.
 
 If multiple calls come through while waiting for 
the interval to expire then after the time expires all calls with go through at the same time
possibly going over your hard limit. Will add a queue system that will process the the calls as they were called.

Along with this a stdout warning will be added to alert with the queue is larger than one interval can hold. 
This will help you debug if you might need a higher limit api key.



