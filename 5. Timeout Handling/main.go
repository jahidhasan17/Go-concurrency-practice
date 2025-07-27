package main

import (
	"fmt"
	"math/rand"
	"time"
)

func simulateTask(done chan<- string) {
	delay := time.Duration(1 + rand.Intn(3)) * time.Second
	time.Sleep(delay)

	done <- fmt.Sprintf("Task completed after %v", delay)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	done := make(chan string)

	go simulateTask(done)

	select {
	case result := <-done:
		fmt.Println(result)
	case <-time.After(2 * time.Second):
		fmt.Println("Time out occurred after 2 second")
	}
}