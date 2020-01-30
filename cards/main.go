package main

import "fmt"

func main() {
	// var card string = "Ace of Spades" - below is different syntax but equivalent
	card := newCard()

	// array is fixed length ; slice is C# list or JS array
	// both must be defined with a data type

	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades") //append returns a new slice

	for i, card := range cards { //range takes each index of the whole slice
		fmt.Println(i, card)
	}

}

// label the data type you expect to return from the function

func newCard() string {
	return "Five of Diamonds"
}
