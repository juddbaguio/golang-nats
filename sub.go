package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/nats.go"
)

type HelpBody struct {
	Client  int         `json:"client"`
	Request interface{} `json:"request"`
	Iter    int         `json:"iter"`
}

func main() {
	log.Println("Starting NATS server")
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println("successfully connected to the NATS broker")

	nc.Subscribe("waiting.request", func(msg *nats.Msg) {
		log.Println("received the following message: ", string(msg.Data))
	})

	nc.Subscribe("help", func(msg *nats.Msg) {
		var payload HelpBody

		json.Unmarshal(msg.Data, &payload)
		log.Printf("Request from Client-%v\n", payload.Client)
		log.Println(payload.Request)
		log.Println("----------------------------------------")

		msg.Respond([]byte(fmt.Sprintf("this is your %v request times, Client-%v", payload.Iter, payload.Client)))
		// nc.Publish(msg.Reply, )
	})
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGABRT)

	<-ch
	nc.Drain()
}
