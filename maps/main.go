package main

import "fmt"

func main() {
	// var colors map[string]string

	colors := make(map[int]string)

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }

	colors[9] = "whitecolor"
	colors[2] = "whitecolor2"

	delete(colors, 2)

	fmt.Println(colors)
}
