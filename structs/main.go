package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	myPerson := person{
		firstName: "Amr",
		lastName:  "Azh",
		contact: contactInfo{
			email:   "amr@gmail.com",
			zipCode: 995533,
		},
	}

	fmt.Printf("%+v", myPerson)
}
