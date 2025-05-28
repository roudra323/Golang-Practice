package main

import (
	"fmt"
	"sync"
)

var (
	data = make(map[string]string)
	rwmu sync.RWMutex
	wg   sync.WaitGroup
)

func write(key, value string) {
	defer wg.Done()
	rwmu.Lock()
	defer rwmu.Unlock()
	fmt.Printf("Writing %s = %s\n", key, value)
	data[key] = value
}

func read(key string) {
	defer wg.Done()
	rwmu.Lock()
	defer rwmu.Unlock()
	fmt.Printf("Reading %s = %s\n", key, data[key])

}

func main() {
	fmt.Println("Testing the RWMutex")

	fooWritten := make(chan struct{})
	bazWritten := make(chan struct{})

	wg.Add(1)
	go func() {
		write("foo", "bar")
		close(fooWritten) // signal write done
	}()

	wg.Add(1)
	go func() {
		<-fooWritten
		read("foo")
	}()

	wg.Add(1)
	go func() {
		<-fooWritten
		read("foo")
	}()

	wg.Add(1)
	go func() {
		write("baz", "qux")
		close(bazWritten)
	}()

	wg.Add(1)
	go func() {
		<-bazWritten
		read("baz")
	}()

	wg.Wait()
}
