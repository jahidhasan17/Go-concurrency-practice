package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(workerId int, ch <-chan int, wg *sync.WaitGroup)  {

	defer wg.Done()

	for task := range ch{
		fmt.Printf("Worker - %v Processing Task : %v\n", workerId, task)
		time.Sleep(time.Second)
		fmt.Printf("Worker - %v Completed Task : %v\n", workerId, task)
	}

	fmt.Printf("Worker - %v Finished it Work\n", workerId)
}

func main()  {
	numWorker := 3
	numTasks := 10

	var wg sync.WaitGroup

	var ch chan int = make(chan int)

	for i := 0; i < numWorker; i++{
		wg.Add(1)
		go worker(i, ch, &wg)
	}

	for i := 0; i < numTasks; i++{
		ch <- i
	}
	close(ch)

	wg.Wait()
	fmt.Println("All Task Finished")
}