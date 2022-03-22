package main

func main() {
	// var card string = "Ace of Spades"
	cards := deck{"Five of Diamonds", "Ten of Diamonds", "One of Diamonds"}
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
