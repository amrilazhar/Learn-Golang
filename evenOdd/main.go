package main

import "fmt"

func main() {
	totalNum := 10
	numbers := []int{}
	for i := 0; i <= totalNum; i++ {
		numbers = append(numbers, i)
	}

	for _, num := range numbers {
		if num%2 == 0 {
			fmt.Println(num, " is even")
		} else {
			fmt.Println(num, " is odd")
		}
	}
}
