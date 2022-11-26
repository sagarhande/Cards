package main

import (
	"fmt"
)

func main() {

	// cards := newDeck()
	// cards.saveToFile("cards.txt")

	newDeckFromFile := newDeckFromFile("cards.txt")
	fmt.Println("Original:  : ", newDeckFromFile)
	newDeckFromFile.shuffle()
	fmt.Println("Shuffle : ", newDeckFromFile)

}
