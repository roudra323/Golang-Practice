/*

Task 1: Basic Goroutines and Channels
You're working at a crypto exchange and need to build a simple price ticker that fetches cryptocurrency prices from multiple sources.
Write a Go program that:

Creates 3 goroutines, each simulating fetching a Bitcoin price from different exchanges (you can use random prices between 40000-50000)
Each goroutine should take a random time between 1-3 seconds to "fetch" the price
Use channels to collect all 3 prices
Print each price as it arrives, along with which exchange it came from


*/

package main

import (
	"fmt"
	"time"
)

func fetchPrice(exchange string, price int, delay time.Duration, ch chan<- string) {
	time.Sleep(delay)
	result := fmt.Sprintf("Exchange %s: $%d", exchange, price)
	ch <- result
}

func main() {
	fmt.Println("Task 1: Basic Goroutines and Channels")

	ch := make(chan string, 3)

	// Launch 3 concurrent goroutines
	go fetchPrice("Binance", 40000, 1*time.Second, ch)
	go fetchPrice("Coinbase", 40500, 2*time.Second, ch)
	go fetchPrice("Kraken", 50000, 3*time.Second, ch)

	// Collect results as they arrive
	for i := 0; i < 3; i++ {
		price := <-ch
		fmt.Printf("Received: %s\n", price)
	}
}
