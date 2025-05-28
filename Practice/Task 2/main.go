/*

Task 2: Building on Task 1 - Adding Context and Cancellation
Now that you understand basic goroutines, let's make our price fetcher more robust.
In a real exchange, you might need to cancel price fetching if it takes too long.
Modify your solution to:

Use context.WithTimeout() to set a 2.5-second timeout for all price fetching
If any exchange takes longer than the timeout, cancel all remaining operations
Handle the context cancellation properly in your goroutines
Print which exchanges succeeded and which were cancelled


*/

package main

import (
	"context"
	"fmt"
	"time"
)

func getPrice(cntx context.Context, messageChannel chan<- string, exchange string, price int, delay time.Duration) {
	timer := time.NewTimer(delay)
	defer timer.Stop()

	select {
	case <-timer.C:
		msgString := fmt.Sprintf("Success - Exchange %s: $%d", exchange, price)

		select {
		case messageChannel <- msgString:
		case <-cntx.Done():
			fmt.Printf("CANCELLED while sending - Exchange %s (context: %v)\n", exchange, cntx.Err())
		}

	case <-cntx.Done():

	}
}

func main() {
	fmt.Println("Task 2: Building on Task 1 - Adding Context and Cancellation")

	cont, cancle := context.WithTimeout(context.Background(), time.Millisecond*300)
	defer cancle()

	messageChannel := make(chan string)

	go func() {
		getPrice(cont, messageChannel, "Binance", 4000, time.Millisecond*100)
		close(messageChannel) // Close when done sending
	}()

	for msg := range messageChannel {
		fmt.Println(msg)
	}

}
