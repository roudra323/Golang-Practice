package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func doWork() int {
	time.Sleep(time.Second)
	return rand.Intn(1000)
}

func experiment1() {
	ch := make(chan int)

	go func() {
		var wg sync.WaitGroup

		for i := 1; i <= 10000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				result := doWork()
				ch <- result

			}()

		}
		wg.Wait()
		close(ch)
	}()

	for n := range ch {
		fmt.Println("Channel Output ", n)
	}
}

func experiment2() {
	ch := make(chan string)

	go func() {
		ch <- "Hello from channel"
	}()

	msg := <-ch

	fmt.Println("Received: ", msg)
}

func experiment3() {
	ch := make(chan int, 3)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3

		close(ch)
	}()

	for i := range ch {
		fmt.Println(i)
	}
}

func main() {
	// experiment1()
	// experiment2()
	experiment3()
}
