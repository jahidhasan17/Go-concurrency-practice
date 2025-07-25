package main

import(
	"fmt"
)

func generateNumber(ch chan<- int)  {
	for i := 1; i <= 100; i++{
		ch <- i
	}

	close(ch)
}

func calculateSum(ch <-chan int, sumCh chan<- int)  {
	sum := 0

	for num := range ch{
		fmt.Println("Current Num is ", num, " And ch size is ", len(ch))
		sum += num
	}

	sumCh <- sum
	close(sumCh)
}

func main()  {
	var ch chan int = make(chan int)
	var sumCh chan int = make(chan int)

	go generateNumber(ch)
	go calculateSum(ch, sumCh)

	fmt.Println("Sum is : ", <-sumCh)
}