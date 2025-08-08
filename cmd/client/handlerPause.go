package main

import (
	"fmt"

	"github.com/SoulOppen/learn-pub-sub-starter/internal/gamelogic"
	"github.com/SoulOppen/learn-pub-sub-starter/internal/routing"
)

func HandlerPause(gs *gamelogic.GameState) func(routing.PlayingState) {
	return func(ps routing.PlayingState) {
		defer fmt.Print("> ")
		gs.HandlePause(ps)

	}
}
