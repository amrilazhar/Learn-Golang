package main

func main() {
	// var card string = "Ace of Spades"
	cards := newDeck()

	hand, remainDeck := deal(cards, 5)
	hand.print()
	remainDeck.print()
}
