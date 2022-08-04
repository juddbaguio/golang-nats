package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	log.Println("Starting NATS publisher")
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println("successfully connected to the NATS broker")

	iter := 0

	clientNumPtr := flag.Int("client", 1, "set your client")
	flag.Parse()

	payloadMap := make(map[string]interface{})

	payloadMap["client"] = *clientNumPtr
	payloadMap["iter"] = iter
	for {
		payloadMap["request"] = fmt.Sprintf("this is my %v request", payloadMap["iter"])
		encodedJson, _ := json.Marshal(payloadMap)

		res, err := nc.Request("help", encodedJson, time.Duration(15)*time.Second)

		if err != nil {
			log.Println(err)
			nc.Close()
			break
		}

		log.Printf("received response: %v\n", string(res.Data))

		time.Sleep(time.Duration(1) * time.Second)
		payloadMap["iter"] = payloadMap["iter"].(int) + 1
	}

}
