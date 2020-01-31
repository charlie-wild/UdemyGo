package main

import "fmt"

/* Create a new type of deck which is a slice of strings */

type deck []string

func (d deck) print() { // any variable of type deck gets access to this print method using the receiver
	for i, card := range d {
		fmt.Println(i, card)
	}
}
