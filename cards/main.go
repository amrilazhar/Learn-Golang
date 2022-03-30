package main

import (
	"cards/decks"
)

func main() {
	// var card string = "Ace of Spades"
	cards := decks.NewDeck()
	// cards.saveToFile("test_deck")

	// cards := newDeckFromFile("test_deck")
	cards.Shuffle()
	cards.Print()
}
