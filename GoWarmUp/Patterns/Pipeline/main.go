package main

import "fmt"

type Car struct {
	name   string
	washed bool
	dry    bool
}

func inputCars(carlist []Car) <-chan Car {
	out := make(chan Car)

	go func() {
		for _, car := range carlist {
			out <- car
		}
		close(out)
	}()

	return out
}

func washedCars(washChan <-chan Car) <-chan Car {
	out := make(chan Car)

	go func() {
		for car := range washChan {
			car.washed = true
			out <- car
		}
		close(out)
	}()

	return out
}

func driedCars(washChan <-chan Car) <-chan Car {
	out := make(chan Car)

	go func() {
		for car := range washChan {
			car.dry = true
			out <- car
		}
		close(out)
	}()

	return out
}

func main() {
	fmt.Printf("Implementing Pipelines\n")

	// slice of cars
	carlist := []Car{
		{name: "Car1"},
		{name: "Car2"},
		{name: "Car3"},
		{name: "Car4"},
		{name: "Car5"},
	}

	// Passing carlist to initical pipeline

	initialPipeline := inputCars(carlist)
	washedCarsPipeline := washedCars(initialPipeline)
	dryCarsPipeline := driedCars(washedCarsPipeline)

	for finalCars := range dryCarsPipeline {
		fmt.Printf("Cars status %+v\n", finalCars)
	}

}
