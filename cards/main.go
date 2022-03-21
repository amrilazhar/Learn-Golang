package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	card := newCard()

	fmt.Println("Let's Play")
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
