package main

import "fmt"

type Number interface {
	int | float32 | int64 | float64
}

func addTwo[T Number](s []T) (result T) {
	for _, numbers := range s {
		result += numbers
	}
	return
}

func main() {
	fmt.Println("Generics")
	arr := [3]int{1, 2, 3}
	result := addTwo(arr[:]) // Convert array to slice with arr[:]
	fmt.Println("Sum:", result)
}
