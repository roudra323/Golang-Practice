package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("We are implementing the fan-in fan-out by our own (havent watched any videos)")

	lNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	jobChan := make(chan int)

	go func() {
		for _, num := range lNumbers {
			jobChan <- num
		}
		close(jobChan)
	}()

	resChan := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for num := range jobChan {
				resChan <- num * num
			}
		}()
	}

	go func() {
		wg.Wait()
		close(resChan)
	}()

	for res := range resChan {
		fmt.Println(res)
	}
}
