package producer

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

// var (
// 	topic         = os.Getenv("KAFKA_TOPIC")
// 	brokerAddress = os.Getenv("KAFKA_HOST") + ":" + os.Getenv("KAFKA_PORT")
// )

const (
	topic = "messages"
	// brokerAddress = {"kafka-1:9092","kafka-2:9092","kafka-3:9092"}
	brokerAddress = "kafka-1:9092"
)

var (
	KafkaCtx    context.Context
	KafkaWriter *kafka.Writer
)

func getKafkaWriter() *kafka.Writer {
	return KafkaWriter
}

func getKafkaCtx() context.Context {
	return KafkaCtx
}

func ProduceMessage(k, val string) error {
	w := getKafkaWriter()
	ctx := getKafkaCtx()

	err := w.WriteMessages(ctx, kafka.Message{
		Key: []byte(k),
		// create an arbitrary message payload for the value
		Value: []byte(val),
		Time:  time.Now(),
	})
	if err != nil {
		fmt.Println("could not write message " + err.Error())
		return err
	}
	return nil
}

func SetupKafka() {
	KafkaCtx = context.Background()

	l := log.New(os.Stdout, "kafka writer: ", 0)

	KafkaWriter = kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		// assign the logger to the writer
		Logger: l,
	})
}
