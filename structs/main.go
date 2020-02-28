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
	//jimPointer := &jim //the & gives us access to the RAM address where the variable is sitting


	 // we don't need the above line as if we have a variable of type person, and the receiver function acts on the pointer, go will turn the variable into a pointer
	 
	 // we don't need to coax out the memory address!
	
	jim.updateName("Jimmy")
	jim.print()
}

// if we see a * where a type should be (as ln 39) it is a type description and not an operator as * is inside the function!

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName

	// the * says to give you the value at the memory address of the variable
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
