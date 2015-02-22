package cards

import (
	"fmt"
	"math/rand"
)

type Card struct {
	Suit byte
	Rank uint8
}

type Deck struct {
	Cards []Card
}

func GenerateStandardDeck() Deck {
	d := Deck{}
	for i := byte(1); i <= 13; i++ {
		for _, suit := range []byte{'D', 'C', 'H', 'S'} {
			card := Card{suit, i}
			d.Cards = append(d.Cards, card)
		}
	}
	return d
}

func (c Card) print() {
	switch {
	case c.Rank == 1:
		fmt.Printf("%c", 'A')
	case c.Rank > 1 && c.Rank < 10:
		fmt.Print(c.Rank)
	case c.Rank == 10:
		fmt.Printf("%c", 'T')
	case c.Rank == 11:
		fmt.Printf("%c", 'J')
	case c.Rank == 12:
		fmt.Printf("%c", 'Q')
	case c.Rank == 13:
		fmt.Printf("%c", 'K')
	}
	fmt.Printf("%c", c.Suit)
}

func (d Deck) Print() {
	for _, card := range d.Cards {
		card.print()
		fmt.Print(" ")
	}
	fmt.Println()
}

func (d Deck) Shuffle() {
	for i := range d.Cards {
		j := rand.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}
