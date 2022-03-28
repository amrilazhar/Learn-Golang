package main

import "fmt"

type bot interface {
	getGreeting() string
	// getFarewell() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	// call getgreeting
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// imagine there is complex code here
	return "Hello there!"
}

func (eb englishBot) getFarewell() string {
	// imagine there is complex code here
	return "Goodbye"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
