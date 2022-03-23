package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// buat type baru namanya deck tipenya slice of string
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Club"}
	cardValue := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// fungsi receiver, mirip implementasi dr inheritance,
// dgn ini type deck bisa punya method bawaan berupa print()
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}
	ss := strings.Split(string(bs), ",")
	return deck(ss)
}
