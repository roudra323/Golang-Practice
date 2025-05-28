package main

import (
	"fmt"
	"sync"
)

var mx sync.Mutex

var count int

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	mx.Lock()
	defer mx.Unlock()
	count++
}

func value() int {
	mx.Lock()
	defer mx.Unlock()
	return count
}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	fmt.Println("The value is : ", value())

}
