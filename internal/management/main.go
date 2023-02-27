package main

import (
	"github.com/nats-io/nats.go"
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

	js.AddStream(&nats.StreamConfig{
		Name:     "myOrderStream",
		Subjects: []string{"ORDER.created"},
	})
}
