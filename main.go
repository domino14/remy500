package main

import "github.com/domino14/remy500/cards"

func main() {
	deck := cards.GenerateStandardDeck()
	deck.Shuffle()
	deck.Print()
}
