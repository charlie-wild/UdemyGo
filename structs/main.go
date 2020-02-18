package main

import "fmt"

//first define fields

type person struct {
	firstName string
	lastName string
}

func main() {
	//alex := person{"Alex","Test"} or
	//alex := person{firstName:"Alex",lastName:"Test"}
	var alex person
	alex.firstName = "Alex"
	fmt.Println(alex)
	fmt.Printf("%+v", alex)
	
}