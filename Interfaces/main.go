package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}


func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

// for receiver functions, we can eliminate the variable and only pass in the type, if we don't use it

func (englishBot) getGreeting() string {
	// super customer logic goes here
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}