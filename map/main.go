package main

import "fmt"

func main() {

	// colors := make(map[int]string) alternative syntax

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	colors["white"] = "#ffffff"

	// no dot syntax available with maps

	delete(colors, "white")

	fmt.Println(colors)
	printMap(colors)

}

func printMap(c map[string]string) {
	for color, key := range c {
		fmt.Println(color, key)
	}
}
