package main

import "fmt"

func greet(name string) {
	fmt.Println("Inside the greet function")
	fmt.Println("Hi,", name)
}

func variablesAndTypes() {
	// Different ways to declare variables
	var name string = "Alice"
	var age int = 25
	var height float64 = 5.6

	// Short declaration (most common)
	city := "New York"
	isStudent := true

	fmt.Printf("Name: %s, Age: %d, Height: %.1f\n", name, age, height)
	fmt.Printf("City: %s, Student: %t\n", city, isStudent)
}

func takeinputs() (float64, float64) {
	var a float64
	var b float64

	fmt.Println("Input the first number: ")

	_, err := fmt.Scan(&a)

	if err != nil {
		fmt.Println("Some Error Happened..")
		return 0, 0 // Return zero values for both float64 returns
	}

	fmt.Println("Input the second number: ")

	_, err = fmt.Scan(&b)

	if err != nil {
		fmt.Println("Some Error Happened..")
		return 0, 0 // Return zero values for both float64 returns
	}

	return a, b // Return the actual values
}

func doSum() {
	num1, num2 := takeinputs()
	fmt.Printf("You entered: %.2f and %.2f\n", num1, num2)
	sum := num1 + num2
	fmt.Printf("The sum is : %.2f\n", sum)
}

func main() {
	greet("Asir")
	variablesAndTypes()
	doSum()
}
