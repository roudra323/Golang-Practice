package main

import "fmt"

// multiplicationsOf2 returns a closure function that multiplies each sequential number by 2
// This is a "factory function" that creates and returns a new function
func multiplicationsOf2() func() int {
	// initial is declared inside the outer function but outside the inner function
	// This variable is "captured" by the closure - it persists between function calls
	// This is what makes closures powerful - they "remember" their environment
	initial := 1

	// The returned anonymous function forms a closure because it "closes over" the initial variable
	return func() int {
		// Each time this function is called, it:
		// 1. Captures the current value of initial
		re := initial
		// 2. Increments the captured variable for the next call
		initial++
		// 3. Returns the current value multiplied by 2
		return re * 2
	}
}

func main() {
	fmt.Println("Learning Closure")

	// two is now a function that maintains its own internal state (initial)
	// It "remembers" the value of initial between calls
	two := multiplicationsOf2()

	// Each call to two() increments its captured initial variable
	// and returns that value multiplied by 2
	for i := 1; i <= 10; i++ {
		fmt.Println(i, " * ", "2 = ", two())
	}

	// If we created another function with multiplicationsOf2(),
	// it would have its own separate initial variable starting at 1
}
