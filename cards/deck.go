package main

import "fmt"

// buat type baru namanya deck tipenya slice of string
type deck []string

// fungsi receiver, mirip implementasi dr inheritance,
// dgn ini type deck bisa punya method bawaan berupa print()
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
