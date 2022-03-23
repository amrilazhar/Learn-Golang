package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck) != 52 {
		t.Errorf("Expected 52 cards, but got %v", len(deck))
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected [Ace of Spades] card, but got %v", deck[0])
	}

	if deck[len(deck)-1] != "King of Clubs" {
		t.Errorf("Expected [King of Clubs] card, but got %v", deck[len(deck)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards, but got %v", len(deck))
	}

	os.Remove("_decktesting")
}
