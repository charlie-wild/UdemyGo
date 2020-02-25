package main

//first define fields
import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo // this is the same as contactInfo (fieldName) contactInfo(struct)
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
		contactInfo: contactInfo{
			email: "email@gmail.com",
			zipCode: 90210, //remember the commas...
		},
	}
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
