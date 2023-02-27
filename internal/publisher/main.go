package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"time"
)

func main() {
	// Connect to NATS
	nc, err := nats.Connect("nats://0.0.0.0:57797", func(options *nats.Options) error {
		options.User = "local"
		options.Password = "eRaESgQ5fBOdDjRRTROgmowzRCYA0P9W"
		return nil
	})

	if err != nil {
		panic(err)
	}

	// Create JetStream Context
	js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		panic(err)
	}

	// Simple Async Stream Publisher
	for i := 0; i < 10000; i++ {
		msg := fmt.Sprintf("[%d] hello", i)
		js.Publish("ORDER.created", []byte(msg))
		time.Sleep(1 * time.Second)
		fmt.Println(msg)
	}
}
