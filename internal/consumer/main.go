package main

import (
	"context"
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

	// Create Pull based consumer with maximum 128 inflight.
	sub, err := js.PullSubscribe("ORDER.created", "wq")
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Timed out after 10 seconds")
			return
		default:
		}

		// Fetch will return as soon as any message is available rather than wait until the full batch size is available, using a batch size of more than 1 allows for higher throughput when needed.
		msgs, _ := sub.Fetch(10, nats.Context(ctx))
		for _, msg := range msgs {
			fmt.Println(string(msg.Data))
			msg.Ack()
		}
	}

	if err := sub.Drain(); err != nil {
		panic(err)
	}
}
