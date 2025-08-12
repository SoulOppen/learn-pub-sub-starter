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
	connection := "amqp://guest:guest@localhost:5672/"
	c, err := amqp.Dial(connection)
	if err != nil {
		log.Fatal("conecction error")
	}
	defer c.Close()
	log.Printf("Connection successful ✅ ")
	cha, err := c.Channel()
	if err != nil {
		log.Fatal("channel error")
	}
	err = pubsub.SubscribeGob(
		c,
		routing.ExchangePerilTopic,
		routing.GameLogSlug,
		routing.GameLogSlug+".*",
		pubsub.Durable,
		handlerLogs(),
	)
	if err != nil {
		log.Fatalf("could not starting consuming logs: %v", err)
	}

	gamelogic.PrintServerHelp()
	for {
		words := gamelogic.GetInput()
		if len(words) == 0 {
			continue
		}
		switch words[0] {
		case "pause":
			fmt.Println("Publishing paused game state")
			err = pubsub.PublishJSON(
				cha,
				routing.ExchangePerilDirect,
				routing.PauseKey,
				routing.PlayingState{
					IsPaused: true,
				},
			)
			if err != nil {
				log.Printf("could not publish time: %v", err)
			}
		case "resume":
			fmt.Println("Publishing resumes game state")
			err = pubsub.PublishJSON(
				cha,
				routing.ExchangePerilDirect,
				routing.PauseKey,
				routing.PlayingState{
					IsPaused: false,
				},
			)
			if err != nil {
				log.Printf("could not publish time: %v", err)
			}
		case "quit":
			log.Println("goodbye")
			return
		default:
			log.Println("unknown command")
		}
	}
}
