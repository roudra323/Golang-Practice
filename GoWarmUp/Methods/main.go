package main

import "fmt"

type Person struct {
	name    string
	age     int
	address string
}

func (p *Person) sayHi() {
	fmt.Printf("The type of p object is %T\n", p)
	fmt.Println("Hi there. I am ", p.name)
}

func (p *Person) changeBirthday() {
	p.age++
}

func main() {
	fmt.Println("Learning Methods")

	var p1 Person = Person{"Roudra", 26, "Bajitpur"}

	fmt.Printf("The type of p1 object is %T\n", p1)

	p1.sayHi()

	p1.changeBirthday()

	fmt.Println("Brithday = ", p1.age)
}
