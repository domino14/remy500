package cards

import (
	"bytes"
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

func (c Card) String() string {
	switch {
	case c.Rank == 1:
		return fmt.Sprintf("A%c", c.Suit)
	case c.Rank > 1 && c.Rank < 10:
		return fmt.Sprintf("%d%c", c.Rank, c.Suit)
	case c.Rank == 10:
		return fmt.Sprintf("T%c", c.Suit)
	case c.Rank == 11:
		return fmt.Sprintf("J%c", c.Suit)
	case c.Rank == 12:
		return fmt.Sprintf("Q%c", c.Suit)
	case c.Rank == 13:
		return fmt.Sprintf("K%c", c.Suit)
	}
	return ""
}

func (d Deck) String() string {
	var buffer bytes.Buffer
	for _, card := range d.Cards {
		buffer.WriteString(card.String())
		buffer.WriteString(" ")
	}
	return buffer.String()
}

func (d Deck) Shuffle() {
	for i := range d.Cards {
		j := rand.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}

// DealCard deals a card to the receiving deck (r).
func (d *Deck) DealCard() Card {
	var card Card
	card, d.Cards = d.Cards[len(d.Cards)-1], d.Cards[:len(d.Cards)-1]
	return card
}

func (d *Deck) AddCard(card Card) {
	d.Cards = append(d.Cards, card)
}
