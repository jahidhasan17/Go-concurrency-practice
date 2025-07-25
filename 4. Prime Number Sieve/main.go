package main

import (
	"fmt"
)

func generate(n int) <-chan int {
	var out chan int = make(chan int)

	go func ()  {
		for i := 2; i <= n; i++{
			out <- i
		}
		close(out)
	}()

	return out
}

func filter(ch <-chan int, prime int) <-chan int {
	var filtered chan int = make(chan int)

	go func ()  {
		for num := range ch{
			if num % prime != 0{
				filtered <- num
			}
		}
		close(filtered)
	}()

	return filtered
}

func main()  {
	ch := generate(100)

	for {
		prime, ok := <- ch

		if !ok {
			break
		}
		
		fmt.Println(prime)

		ch = filter(ch, prime)
	}
}
