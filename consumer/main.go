package main

// Separate application for kafka brokers, read messagees produces by main app

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

const (
	topic = "messages"
	// brokerAddress = [3]string{"kafka-1:9092", "kafka-2:9092", "kafka-3:9092"}
	// brokerAddress = "kafka-1:9092"
)

func getBrokers() []string {
	return []string{"kafka-1:9092", "kafka-2:9092", "kafka-3:9092"}
}

func main() {
	// create a new context
	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking
	consume(ctx)
}

func consume(ctx context.Context) {
	// create a new logger that outputs to stdout
	// and has the `kafka reader` prefix
	// l := log.New(os.Stdout, "kafka reader: ", 0)
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafka.NewReader(kafka.ReaderConfig{
		// Brokers: []string{brokerAddress},
		Brokers: getBrokers(),
		Topic:   topic,
		GroupID: "group1",
		// assign the logger to the reader
		// Logger: l,
	})
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		fmt.Println("received: ", string(msg.Value), msg.Time)
	}
}
