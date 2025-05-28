package main

import (
	"fmt"
	"time"
)

func doWork(ch <-chan bool) {

	for {
		select {
		case <-ch:
			return
		default:
			fmt.Println("Working")
		}
	}

}

func main() {
	done := make(chan bool)

	go doWork(done)

	time.Sleep(time.Millisecond * 5)

	close(done)

}
