package main

import "fmt"

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)

	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()

	return out
}

func squareNumbers(chn <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for num := range chn {
			out <- num * num
		}
		close(out)
	}()

	return out
}

func main() {

	// input
	nums := []int{2, 3, 4, 7, 1}

	// Stage 1
	dataChannel := sliceToChannel(nums)

	// Stage 2
	finalChannel := squareNumbers(dataChannel)

	// Stage 3
	for n := range finalChannel {
		fmt.Println(n)
	}

}
