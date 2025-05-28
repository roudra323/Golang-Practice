package main

import "fmt"

func doTask(ch chan<- bool) {
	fmt.Println("Doing task")

	ch <- true

	close(ch)
}

func main() {
	donech := make(chan bool)

	go doTask(donech)

	select {
	case <-donech:
		fmt.Println("Do task is executed")
	}
}
