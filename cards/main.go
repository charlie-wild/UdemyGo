package main

func main() {
	// var card string = "Ace of Spades" - below is different syntax but equivalent
	//card := newCard()

	// array is fixed length ; slice is C# list or JS array
	// both must be defined with a data type

	cards := newDeck()
	//cards = append(cards, "Six of Spades") //append returns a new slice
	//cards.print()

	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()

}

// label the data type you expect to return from the function

func newCard() string {
	return "Five of Diamonds"
}
