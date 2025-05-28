package main

import "fmt"

// Generics in Go (added in Go 1.18) allow you to write functions and data structures
// that can work with any data type that satisfies certain constraints.

// PrintSlice is a generic function that can print any slice of values
// The type parameter 'T' can be any type
func PrintSlice[T any](s []T) {
	fmt.Print("[ ")
	for _, v := range s {
		fmt.Printf("%v ", v)
	}
	fmt.Println("]")
}

// Map is a generic function that applies a function to each element in a slice
// and returns a new slice with the results
func Map[T, U any](s []T, f func(T) U) []U {
	result := make([]U, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

// Comparable constraint interface that allows only types that can be compared with == and !=
// FindIndex finds the index of an element in a slice, or -1 if not found
func FindIndex[T comparable](s []T, item T) int {
	for i, v := range s {
		if v == item {
			return i
		}
	}
	return -1
}

// CustomConstraint demonstrates defining a custom constraint interface
// This constraint only allows numeric types
type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

// Sum adds up all the numbers in a slice
// The constraint ensures we only work with numeric types
func Sum[T Number](s []T) T {
	var result T
	for _, v := range s {
		result += v
	}
	return result
}

// Stack is a generic data structure example
// It can hold elements of any type
type Stack[T any] struct {
	items []T
}

// Push adds an element to the top of the stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top element of the stack
// Returns the zero value of T and false if the stack is empty
func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}

	// Get the last item
	item := s.items[len(s.items)-1]
	// Remove the last item
	s.items = s.items[:len(s.items)-1]
	return item, true
}

// IsEmpty checks if the stack has no elements
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of elements in the stack
func (s *Stack[T]) Size() int {
	return len(s.items)
}

// A more advanced example: Filter function with constraints
func Filter[T any](s []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range s {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	fmt.Println("===== Go Generics Examples =====")

	// Example 1: Using our generic PrintSlice function with various types
	fmt.Println("\n--- Example 1: PrintSlice ---")
	intSlice := []int{1, 2, 3, 4, 5}
	stringSlice := []string{"apple", "banana", "cherry"}
	floatSlice := []float64{1.1, 2.2, 3.3}

	fmt.Print("Integers: ")
	PrintSlice(intSlice)
	fmt.Print("Strings: ")
	PrintSlice(stringSlice)
	fmt.Print("Floats: ")
	PrintSlice(floatSlice)

	// Example 2: Using our Map function to transform slices
	fmt.Println("\n--- Example 2: Map ---")
	// Double each integer
	doubled := Map(intSlice, func(x int) int {
		return x * 2
	})
	fmt.Print("Original: ")
	PrintSlice(intSlice)
	fmt.Print("Doubled: ")
	PrintSlice(doubled)

	// Convert integers to strings
	strNumbers := Map(intSlice, func(x int) string {
		return fmt.Sprintf("Number %d", x)
	})
	fmt.Print("Strings from ints: ")
	PrintSlice(strNumbers)

	// Example 3: Using FindIndex with different types
	fmt.Println("\n--- Example 3: FindIndex ---")
	fmt.Printf("Index of 3 in intSlice: %d\n", FindIndex(intSlice, 3))
	fmt.Printf("Index of 'banana' in stringSlice: %d\n", FindIndex(stringSlice, "banana"))
	fmt.Printf("Index of 'orange' in stringSlice: %d\n", FindIndex(stringSlice, "orange"))

	// Example 4: Using Sum with the Number constraint
	fmt.Println("\n--- Example 4: Sum with constraint ---")
	fmt.Printf("Sum of integers: %d\n", Sum(intSlice))
	fmt.Printf("Sum of floats: %.2f\n", Sum(floatSlice))
	// The line below would not compile:
	// Sum(stringSlice) // Error: string does not satisfy Number constraint

	// Example 5: Using our generic Stack
	fmt.Println("\n--- Example 5: Stack data structure ---")
	// Integer stack
	intStack := Stack[int]{}
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)
	fmt.Printf("Int Stack size: %d\n", intStack.Size())

	if val, ok := intStack.Pop(); ok {
		fmt.Printf("Popped: %d\n", val)
	}
	if val, ok := intStack.Pop(); ok {
		fmt.Printf("Popped: %d\n", val)
	}
	fmt.Printf("Stack size after pops: %d\n", intStack.Size())

	// String stack
	stringStack := Stack[string]{}
	stringStack.Push("Go")
	stringStack.Push("is")
	stringStack.Push("awesome")

	fmt.Print("String Stack contents: ")
	for !stringStack.IsEmpty() {
		val, _ := stringStack.Pop()
		fmt.Printf("%s ", val)
	}
	fmt.Println()

	// Example 6: Advanced filter usage
	fmt.Println("\n--- Example 6: Filter ---")
	// Filter even numbers
	evenNumbers := Filter(intSlice, func(x int) bool {
		return x%2 == 0
	})
	fmt.Print("Even numbers: ")
	PrintSlice(evenNumbers)

	// Filter strings that start with 'a'
	aWords := Filter(stringSlice, func(s string) bool {
		return len(s) > 0 && s[0] == 'a'
	})
	fmt.Print("Words starting with 'a': ")
	PrintSlice(aWords)

	fmt.Println("\n===== End of Go Generics Examples =====")
}
