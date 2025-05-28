package main

import "fmt"

// Step 1: Basic Hello World
func helloWorld() {
	fmt.Println("Hello, World!")
}

// Step 2: Variables and Types
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

// Step 3: Functions with parameters and return values
func add(a, b int) int {
	return a + b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// Step 4: Arrays and Slices
func arraysAndSlices() {
	// Array (fixed size)
	var numbers [3]int = [3]int{1, 2, 3}
	fmt.Println("Array:", numbers)

	// Slice (dynamic array)
	fruits := []string{"apple", "banana", "orange"}
	fruits = append(fruits, "grape")
	fmt.Println("Slice:", fruits)
	fmt.Println("Length:", len(fruits))
}

// Step 5: Maps (key-value pairs)
func maps() {
	// Creating a map
	ages := make(map[string]int)
	ages["Alice"] = 25
	ages["Bob"] = 30

	// Map literal
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}

	fmt.Println("Ages:", ages)
	fmt.Println("Colors:", colors)

	// Check if key exists
	if age, exists := ages["Alice"]; exists {
		fmt.Printf("Alice is %d years old\n", age)
	}
}

// Step 6: Loops
func loops() {
	// For loop (traditional)
	fmt.Print("Count up: ")
	for i := 1; i <= 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// For loop with slice (range)
	numbers := []int{10, 20, 30, 40}
	fmt.Print("Numbers: ")
	for index, value := range numbers {
		fmt.Printf("[%d]=%d ", index, value)
	}
	fmt.Println()

	// While-style loop
	count := 0
	fmt.Print("While style: ")
	for count < 3 {
		fmt.Print(count, " ")
		count++
	}
	fmt.Println()
}

// Step 7: If statements and conditionals
func conditionals() {
	score := 85

	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	// If with initialization
	if num := 42; num%2 == 0 {
		fmt.Printf("%d is even\n", num)
	} else {
		fmt.Printf("%d is odd\n", num)
	}
}

// Step 8: Structs (custom types)
type Person struct {
	Name string
	Age  int
	City string
}

func structs() {
	// Creating struct instances
	p1 := Person{Name: "Alice", Age: 25, City: "NYC"}
	p2 := Person{"Bob", 30, "LA"} // positional

	fmt.Printf("Person 1: %+v\n", p1)
	fmt.Printf("Person 2: %s is %d years old\n", p2.Name, p2.Age)

	// Struct pointer
	p3 := &Person{Name: "Charlie", Age: 35, City: "Chicago"}
	p3.Age = 36 // Go automatically dereferences
	fmt.Printf("Person 3 Type: %T\n", p3)
	fmt.Printf("Person 3: %+v\n", *p3)
}

// Step 9: Methods on structs
func (p Person) introduce() string {
	return fmt.Sprintf("Hi, I'm %s from %s", p.Name, p.City)
}

func (p *Person) haveBirthday() {
	(*p).Age++
}

func methods() {
	person := Person{Name: "Diana", Age: 28, City: "Boston"}
	fmt.Println(person.introduce())

	person.haveBirthday()
	fmt.Printf("After birthday: %s is now %d\n", person.Name, person.Age)
}

// Step 10: Error handling
func safeDiv(a, b float64) {
	result, err := divide(a, b)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("%.2f / %.2f = %.2f\n", a, b, result)
}

func main() {
	fmt.Println("=== Go Warmup Tutorial ===\n")

	fmt.Println("1. Hello World:")
	helloWorld()

	fmt.Println("\n2. Variables and Types:")
	variablesAndTypes()

	fmt.Println("\n3. Functions:")
	sum := add(10, 5)
	fmt.Printf("10 + 5 = %d\n", sum)

	fmt.Println("\n4. Arrays and Slices:")
	arraysAndSlices()

	fmt.Println("\n5. Maps:")
	maps()

	fmt.Println("\n6. Loops:")
	loops()

	fmt.Println("\n7. Conditionals:")
	conditionals()

	fmt.Println("\n8. Structs:")
	structs()

	fmt.Println("\n9. Methods:")
	methods()

	fmt.Println("\n10. Error Handling:")
	safeDiv(10, 2)
	safeDiv(10, 0)
}
