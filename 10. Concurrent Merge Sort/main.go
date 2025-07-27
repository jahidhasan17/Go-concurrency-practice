package main

import (
	"math/rand"
	"fmt"
)

func mergeSortC(nums []int, out chan<- int) {
	n := len(nums)

	if n == 0{
		close(out)
		return
	}
	if n == 1{
		out <- nums[0]
		close(out)
		return
	}

	mid := n / 2

	left := make(chan int)
	right := make(chan int)

	go mergeSortC(nums[:mid], left)
	go mergeSortC(nums[mid:], right)

	mergeC(left, right, out)
}

func mergeC(left chan int, right chan int, result chan<- int)  {
	defer close(result)
	val1, ok1 := <-left
	val2, ok2 := <-right

	for ok1 && ok2 {
		if val1 < val2{
			result <- val1
			val1, ok1 = <-left
		} else{
			result <- val2
			val2, ok2 = <-right
		}
	}

	for ok1 {
		result <- val1
		val1, ok1 = <-left
	}

	for ok2 {
		result <- val2
		val2, ok2 = <-right
	}
}

func concurrentMergeSort()  {
	data := make([]int, 10)

	for i := range data {
		data[i] = rand.Intn(10)
	}

	fmt.Println("Data Before Sort", data)
	var result chan int = make(chan int)

	go mergeSortC(data, result)

	fmt.Println("Data after Sort")
	for num := range result{
		fmt.Print(num, " ")
	}
}

func main() {
	concurrentMergeSort()
}