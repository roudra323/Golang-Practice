package main

import (
	"fmt"
)

func sendAMessage(ch1 chan string) {
	ch1 <- "This is from channel 1"
}

func sendAnotherMessage(ch2 chan string) {
	ch2 <- "This is from channel 2"
}

func unbufferedChannel() {
	fmt.Println("Preparing Unbuffered Channels....")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go sendAMessage(ch1)
	go sendAnotherMessage(ch2)

	select {
	case msg := <-ch1:
		fmt.Println(msg)

	case msg := <-ch2:
		fmt.Println(msg)
	}
}

func putValueinBuffered(chBuf chan int) {
	for i := 1; i <= 4; i++ {
		chBuf <- i
	}

	close(chBuf)
}

func bufferedChannel() {

	chaBuff := make(chan int, 4)

	putValueinBuffered(chaBuff)

	for value := range chaBuff {
		fmt.Println(value)
	}
}

func main() {
	// unbufferedChannel()
	bufferedChannel()
}
