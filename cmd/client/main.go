package main

import (
	"fmt"
	"log"

	"github.com/SoulOppen/learn-pub-sub-starter/internal/gamelogic"
	"github.com/SoulOppen/learn-pub-sub-starter/internal/pubsub"
	"github.com/SoulOppen/learn-pub-sub-starter/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting Peril client...")
	connection := "amqp://guest:guest@localhost:5672/"
	c, err := amqp.Dial(connection)
	if err != nil {
		log.Fatal("conecction error")
	}
	defer c.Close()
	user, err := gamelogic.ClientWelcome()
	if err != nil {
		log.Fatal("welcome err")
	}
	log.Printf("welcome %s\n", user)

	gs := gamelogic.NewGameState(user)
	err = pubsub.SubscribeJSON(c, routing.ExchangePerilDirect, fmt.Sprintf("%s.%s", routing.PauseKey, user), routing.PauseKey, "transient", HandlerPause(gs))
	if err != nil {
		log.Fatal(err)
	}
	err = pubsub.SubscribeJSON(c, routing.ExchangePerilTopic, fmt.Sprintf("%s.%s", routing.ArmyMovesPrefix, user), fmt.Sprintf("%s.*", routing.ArmyMovesPrefix), "transient", HandlerPause(gs))
	if err != nil {
		log.Fatal(err)
	}
	gamelogic.PrintServerHelp()
	for {
		words := gamelogic.GetInput()
		if len(words) == 0 {
			continue
		}
		switch words[0] {
		case "move":
			_, err := gs.CommandMove(words)
			if err != nil {
				fmt.Println(err)
				continue
			}

			// TODO: publish the move
		case "spawn":
			err = gs.CommandSpawn(words)
			if err != nil {
				fmt.Println(err)
				continue
			}
		case "status":
			gs.CommandStatus()
		case "help":
			gamelogic.PrintClientHelp()
		case "spam":
			// TODO: publish n malicious logs
			fmt.Println("Spamming not allowed yet!")
		case "quit":
			gamelogic.PrintQuit()
			return
		default:
			fmt.Println("unknown command")
		}
	}
}
