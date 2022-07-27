package main

func main() {
	cards := deck{"one", "two", "three", newCard()}
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
