package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type CrawlerResult struct {
	url      string
	body	 string
	duration time.Duration
	error	 error
}

func crawl(urls []string, concurrency int) <-chan CrawlerResult {
	var results chan CrawlerResult = make(chan CrawlerResult)

	var semaphore chan struct{} = make(chan struct{}, concurrency)
	var wg sync.WaitGroup

	for i := 0; i < len(urls); i++{
		wg.Add(1)
		go func (url string)  {
			defer wg.Done()

			// Acquire lock
			semaphore <- struct{}{}
			
			result := fetchUrl(url)

			results <- result
			
			// Release lock
			<-semaphore
		}(urls[i])
	}

	go func ()  {
		wg.Wait()
		close(results)
	}()

	return results
}

func main() {
	urls := []string{
		"site1.com",
		"site2.com",
		"site3.com",
		"site4.com",
		"site5.com",
		"site6.com",
		"site7.com",
		"site8.com",
		"site9.com",
		"site10.com",
		"site11.com",
		"site12.com",
	}

	var concurrency = 4

	results := crawl(urls, concurrency)


	fmt.Println("Printing Result")
	for result := range results{
		if result.error != nil {
			fmt.Println("Found Error :", result.error, "Url is :", result.url, "Duration took", result.duration)
		} else {
			fmt.Println("Fetched Url :", result.url, "Duration Took :", result.duration)
		}
	}
}

func fetchUrl(url string)  CrawlerResult{
	start := time.Now()
	time.Sleep(time.Duration(rand.Intn(900) + 100) * time.Millisecond)

	if rand.Float32() < 0.2{
		return CrawlerResult{
			url: url,
			body: "",
			duration: time.Since(start),
			error: fmt.Errorf("Network Error"),
		}
	}

	return CrawlerResult{
		url: url,
		body: "body",
		duration: time.Since(start),
		error: nil,
	}
}