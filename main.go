package main

func main() {
	cards := newDeckFromFile("mycards")
	cards.shuffleDeck()
	cards.print()
}
