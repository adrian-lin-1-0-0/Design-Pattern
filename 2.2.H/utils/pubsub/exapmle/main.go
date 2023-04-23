package main

import (
	"fmt"

	"2.2.H/utils/pubsub"
)

func main() {
	hub := pubsub.NewHub(nil)

	sub1 := hub.Subscribe("topic-1")
	sub2 := hub.Subscribe("topic-2")

	go func() {
		for {
			select {
			case msg := <-sub1.MessageCh:
				fmt.Printf("Received message on topic-1: %+v\n", msg)
			case msg := <-sub2.MessageCh:
				fmt.Printf("Received message on topic-2: %+v\n", msg)
			}
		}
	}()

	hub.Publish("topic-1", "Hello, topic 1!")
	hub.Publish("topic-2", "Hello, topic 2!")
}
