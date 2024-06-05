package main

import "fmt"

func main() {
	//my_array := [...]string{"1", "2", "3", "4", "5"}
	// v := 10
	// b := 20
	// for i, v := range my_array {
	// 	fmt.Println(i, v)
	// }
	// cards := newDeck()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()
	// area, perimeter := rectangleDimensions(5.0, 3.0)
	// fmt.Println("Area:", area, "Perimeter:", perimeter) // Output: Area: 15
	// //fmt.Println("Perimeter:", perimeter)
	// runExamples()
	message := func(name string) string {
		return "Hello, " + name
	}("Go")
	fmt.Println(message)
}
