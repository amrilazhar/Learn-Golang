package main

import "fmt"

type bot interface {
	getGreeting(string) string
	// getFarewell() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb, "hi")
	printGreeting(sb, "hola")

}

func printGreeting(b bot, msg string) {
	// call getgreeting
	fmt.Println(b.getGreeting("start"))
}

func (eb englishBot) getGreeting(s string) string {
	// imagine there is complex code here
	return "Hello there!"
}

func (eb englishBot) getFarewell() string {
	// imagine there is complex code here
	return "Goodbye"
}

func (sb spanishBot) getGreeting(s string) string {
	return "Hola!"
}
