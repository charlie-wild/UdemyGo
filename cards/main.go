package main

import "fmt"

func main() {
	// var card string = "Ace of Spades" - below is different syntax but equivalent
	card := newCard()
	fmt.Println(card)
}

// label the data type you expect to return from the function

func newCard() string {
	return "Five of Diamonds"
}
