package main

func main() {
	card := deck{newCard(), "b"}
	card = append(card, "new string")
	//fmt.Println(card)
	card.print()
}

func newCard() string {
	return "this is string"
}
