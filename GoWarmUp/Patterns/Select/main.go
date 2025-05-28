package main

import (
	"context"
	"fmt"
	"time"
)

func getData(ch chan<- string, name string, t time.Duration, ctx context.Context) {

	select {
	case <-time.After(t):
		select {
		case ch <- name:
		case <-ctx.Done():
			// Don't send if context already cancelled
		}
	case <-ctx.Done():
		fmt.Println("Timeout")
	}
}

func main() {
	fmt.Println("Using select")

	// Use cancel instead of error for the context cancel function name
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // Use the correct name here

	ch := make(chan string)

	go getData(ch, "Data A", 500*time.Millisecond, ctx)
	go getData(ch, "Data B", 200*time.Millisecond, ctx)

	select {
	case msg := <-ch:
		fmt.Println("Got data from Channel ", msg)
	case <-ctx.Done():
		fmt.Println("Context canceled:", ctx.Err())

	}
}
