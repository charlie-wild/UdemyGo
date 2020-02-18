package main

import "fmt"

//first define fields

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

func main() {
	//alex := person{"Alex","Test"} or
	//alex := person{firstName:"Alex",lastName:"Test"}
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	jim := person{
		firstName: "Jim",
		lastName: "Party",
		contact: contactInfo{
			email: "email@gmail.com",
			zipCode: 90210, //remember the commas...
		},
	}
}