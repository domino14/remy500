package main

import (
	"github.com/domino14/remy500/rummy500"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	players := []rummy500.Player{}
	players = append(players, rummy500.Player{Username: "foo"})
	players = append(players, rummy500.Player{Username: "bar"})

	rummy500.InitGame(players)
}
