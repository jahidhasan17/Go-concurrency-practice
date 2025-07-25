package main

import(
	"fmt"
	"sync"
)

func Producer(ch chan<- int, wg *sync.WaitGroup)  {
	defer wg.Done()
	for i := 1; i <= 5; i++{
		ch <- i
	}

	close(ch)
}

func Receiver(ch <- chan int, wg *sync.WaitGroup){
	defer wg.Done()
	for msg := range ch{
		fmt.Println(msg)
	}
}

func main()  {
	var ch chan int = make(chan int)

	// WaitGroup is beging used here for keeping main thread active until producing and receving work done.
	var wg sync.WaitGroup

	wg.Add(2)

	go Producer(ch, &wg)
	go Receiver(ch, &wg)

	wg.Wait()
}