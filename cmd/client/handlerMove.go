package main

import (
	"fmt"

	"github.com/SoulOppen/learn-pub-sub-starter/internal/gamelogic"
)

func HandlerMove(gs *gamelogic.GameState) func(gamelogic.ArmyMove) {
	return func(ps gamelogic.ArmyMove) {
		defer fmt.Print("> ")
		gs.HandleMove(ps)

	}
}
