package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Shaw",
		contact: contactInfo{
			email:   "jimxshaw@gmail.com",
			zipCode: 90210,
		},
	}

	fmt.Printf("%+v", jim)
}
