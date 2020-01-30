package main

import "fmt"

func main() {
	// var card string = "Ace of Spades" - below is different syntax but equivalent
	card := newCard()
	fmt.Println(card)

	// array is fixed length ; slice is C# list or JS array
	// both must be defined with a data type

	cards := []string{"Ace of Diamonds", newCard()}
}

// label the data type you expect to return from the function

func newCard() string {
	return "Five of Diamonds"
}
