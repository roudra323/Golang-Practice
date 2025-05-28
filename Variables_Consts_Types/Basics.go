package main

import "fmt"

func main() {
	var langName string = "Go"
	var myName string = "Roudra"
	var fullText string = myName + " is learning " + langName

	fmt.Println(fullText)
	extraInfo(myName)
	getRune()
	testingForLoop()
}

func extraInfo(_name string) {
	const Pi = 3.1416
	fmt.Println(_name + " can count " + fmt.Sprint(Pi))
}

func getRune() {
	str := "你好"
	for i, r := range str {
		fmt.Printf("Index %d: rune %c (Unicode %U)\n", i, r, r)
	}

}

func testingForLoop() {

	fmt.Println("Classic for loop")
	// Classic for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// // While-style
	// i := 0
	// for i < 5 {
	// 	i++
	// }

	str := "hi this is roudra"

	for i, r := range str {
		fmt.Println(i, r)
	}

}
