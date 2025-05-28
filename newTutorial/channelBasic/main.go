package main

import "fmt"

func inputToChannel(numbers []int) <-chan int {
	out := make(chan int)

	go func() {
		for _, num := range numbers {
			out <- num
		}

		close(out)
	}()

	return out
}

func squareElements(chn <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for num := range chn {
			out <- num * num
		}
		close(out)
	}()

	return out
}

func qubeElements(chn <-chan int) <-chan int {
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
	numSlice := []int{2, 4, 6, 8, 9, 1}

	numChan := inputToChannel(numSlice)

	squareChan := squareElements(numChan)

	// for res := range squareChan {
	// 	fmt.Println(res)
	// }

	finalQube := qubeElements(squareChan)

	for res := range finalQube {
		fmt.Println(res)
	}
}
