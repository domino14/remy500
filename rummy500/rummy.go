package rummy500

import (
	"fmt"
	"github.com/domino14/remy500/cards"
)

type Player struct {
	Username string
	Hand     cards.Deck
}

// The face-down cards
type StockPile struct {
	cards cards.Deck
}

// The face-up cards
type DiscardPile struct {
	cards cards.Deck
}

func CreatePlayer(username string) Player {
	player := Player{Username: username}
	return player
}

// InitGame initializes the game and deals the cards, etc.
func InitGame(players []Player) {
	var perPlayer int
	numPlayers := len(players)
	if numPlayers == 2 {
		perPlayer = 13
	} else if numPlayers == 3 || numPlayers == 4 {
		perPlayer = 7
	}
	deck := cards.GenerateStandardDeck()
	deck.Shuffle()
	fmt.Println("Generated deck", deck)
	dealCards(players, deck, perPlayer)
	for _, p := range players {
		p.PrintHand()
	}
}

// dealCards deals cards to a set of players.
func dealCards(players []Player, deck cards.Deck, perPlayer int) {
	for i := 0; i < perPlayer; i++ {
		for j := range players {
			card := deck.DealCard()
			players[j].Hand.AddCard(card)
		}
	}
}

func (p Player) PrintHand() {
	fmt.Println("Player", p.Username, "has hand", p.Hand)
}
