package main

import "fmt"

func main() {
	var a [3]int = [3]int{1, 2, 3}

	for i := 0; i < 3; i++ {
		fmt.Println(a[i])

	}

	var s []int = []int{10, 20, 30}
	fmt.Println(s)

	s = append(s, 40)
	fmt.Println(s)

	// slicing
	fmt.Println(s[1:3])

	fmt.Println(len(s), cap(s))

	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}

	ages["Charlie"] = 35

	delete(ages, "Bob")

	fmt.Println(ages)

	value, exists := ages["Bob"]

	if exists {
		fmt.Println("Bob's age: ", value)
	} else {
		fmt.Println("Bob not found")
	}

	testStructs()

	pointers()

}

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Println("Hello, my name is", p.Name)
	p.Name = "Xlice"
	fmt.Println("Hello, my changed name is", p.Name)

}

func testStructs() {
	p := Person{Name: "Alice", Age: 30}
	p.Greet()
	fmt.Println(p.Name) // Alice
	fmt.Println(p.Age)
}

func pointers() {

	x := 10

	p := &x

	fmt.Println("This is address ", p)
	fmt.Println("This is value ", *p)

}
