package main

import "fmt"

func intDelta(a *int) {
	*a = 52
	fmt.Printf("%v\t%T\n", a, a)
}

func main() {
	a := 42
	fmt.Println("Value before function: ", a)
	intDelta(&a)
	fmt.Println("Value after function: ", a)
}
