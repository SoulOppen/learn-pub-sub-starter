package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

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
	_, _, err = pubsub.DeclareAndBind(c, routing.ExchangePerilDirect, fmt.Sprintf("%s.%s", routing.PauseKey, user), routing.PauseKey, "transient")
	if err != nil {
		log.Fatal("channel err")
	}
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	<-sigChan

	log.Println("\n🔴 Signal received. Shutting down...")
	log.Println("✅ Connection closed. Bye!")

}
