package main

import "fmt"

type bot interface {
	getGreeting() string //if you are a type in this program with a func called getGreeting and return a string then you 
												// are a member of type bot
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {   //members of type bot can call this function 
	fmt.Println(b.getGreeting())
}




// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }



// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

// for receiver functions, we can eliminate the variable and only pass in the type, if we don't use it

func (englishBot) getGreeting() string {
	// super customer logic goes here
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola Chica!"
}