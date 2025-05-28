package main

import "fmt"

var m = make(map[int][]int)

func multipli(val int) {

	for i := val; i <= 100; i += val {
		m[val] = append(m[val], i)
	}
	fmt.Println(m)
}

func main() {
	for i := 2; i <= 10; i++ {
		multipli(i)
	}
}
