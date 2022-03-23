package main

func main() {
	// var card string = "Ace of Spades"
	// cards := newDeck()
	// cards.saveToFile("test_deck")

	cards := newDeckFromFile("test_deck")
	cards.shuffle()
	cards.print()
}
