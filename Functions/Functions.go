package main

import "fmt"

func main() {

	var numerator int
	var denumerator int

	fmt.Println("Enter numerator value: ")

	_, err := fmt.Scan(&numerator)
	if err != nil {
		fmt.Println("Error in input: ", err)
		return
	}

	fmt.Println("Enter denumerator value: ")

	_, err = fmt.Scan(&denumerator)

	if err != nil {
		fmt.Println("Error in input: ", err)
		return
	}

	var result float64 = getFloatValue(numerator, denumerator)
	fmt.Println("This is float value from getFloatValue function-> ", result)
}

func getFloatValue(numerator int, denominator int) float64 {
	if denominator == 0 {
		return 0
	}
	return float64(numerator) / float64(denominator)
}
