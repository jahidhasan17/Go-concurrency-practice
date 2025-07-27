package main

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiter struct {
	tokens chan struct{}
	stop chan struct{}
	wg sync.WaitGroup
}

func NewRateLimiter(limit int, interval time.Duration) *RateLimiter{
	rl := &RateLimiter{
		tokens: make(chan struct{}, limit),
		stop: make(chan struct{}),
	}

	for i := 0; i < limit; i++{
		rl.tokens <- struct{}{}
	}
	rl.wg.Add(1)

	go rl.refillTokens(limit, interval)

	return rl
}

func (rl *RateLimiter) refillTokens(limit int, interval time.Duration)  {
	ticker := time.NewTicker(interval)
	defer rl.wg.Done()
	defer ticker.Stop()
	
	for {
		select{
		case <-ticker.C:
			fmt.Println("Refilling Token and Currenly token size is", len(rl.tokens))
			for i := 0; i < limit; i++{
				select{
					case rl.tokens <- struct{}{}:
					default:
						break
				}
			}
		case <-rl.stop:
			fmt.Println("RateLimiter is Stopping...")
			return
		}
	}
}

func (rl *RateLimiter) Wait() {
	<-rl.tokens
}

func (rl *RateLimiter) Stop(){
	close(rl.tokens)
	close(rl.stop)
	rl.wg.Wait()
}

func main() {
	limiter := NewRateLimiter(5, time.Second * 3)
	defer limiter.Stop()
	
	var wg sync.WaitGroup

	for i := 0; i < 20; i++{
		wg.Add(1)
		time.Sleep(time.Second)
		go func (i int)  {
			defer wg.Done()

			limiter.Wait()
			fmt.Println("Executing Operation", i)
		}(i)
	}

	wg.Wait()
}