package main

import (
	"fmt"
	"sync"
)

func generate(in chan<- int, n int){
	go func ()  {
		for i := 1; i <= n; i++{
			fmt.Println("Generating => ", i);
			in <- i
		}
		close(in)
	}()
}

func square(out <-chan int) <-chan int{
	var in chan int = make(chan int)

	go func ()  {
		for num := range out{
			fmt.Println("Squaring => ", num, num * num);
			in <- num * num
		}
		close(in)
	}()

	return in
}

func merge(out ...<-chan int) <-chan int{
	var mergedCh chan int = make(chan int)
	var wg sync.WaitGroup

	n := len(out)
	wg.Add(n)

	for i := 0; i < n; i++{
		go func (inCh <-chan int)  {
			for num := range inCh{
				fmt.Println("Merging => ", num);
				mergedCh <- num
			}
			wg.Done()
		}(out[i])
	}

	go func ()  {
		wg.Wait()
		close(mergedCh)
	}()

	return mergedCh
}

func main()  {
	var in chan int = make(chan int)

	go generate(in, 10)
	
	ch1 := square(in)
	ch2 := square(in)
	ch3 := square(in)

	for num := range merge(ch1, ch2, ch3){
		fmt.Println(num)
	}
}