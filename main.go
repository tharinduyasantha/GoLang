package main

import "fmt"

func main() {
	card := newCard()
	fmt.Println(card)
	printState()
}

func newCard() string {
	return "this is string"
}
