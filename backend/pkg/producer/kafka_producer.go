// message producer package for kafka

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

type KafkaLogger struct {
	topic           string
	brokerAddresses []string
	// context for creating messages
	kafkaCtx context.Context
	// object on which behalf messages are sent
	kafkaWriter *kafka.Writer
}

type KafkaLoggerInterface interface {
	SetupKafka()
	ProduceMessage(k, val string) error
}

// const (
// 	topic = "messages"
// 	// brokerAddress = {"kafka-1:9092","kafka-2:9092","kafka-3:9092"}
// 	brokerAddress = "kafka-1:9092"
// )

// var ()

// TODO odczyt zmiennych środowiskowych zamiast hardocoded stałych
func (kl *KafkaLogger) SetupKafka() {

	kl.topic = "messages"
	kl.brokerAddresses = []string{"kafka-1:9092", "kafka-2:9092", "kafka-3:9092"}
	kl.kafkaCtx = context.Background()

	l := log.New(os.Stdout, "kafka writer: ", 0)

	kWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers: kl.brokerAddresses,
		Topic:   kl.topic,
		// assign the logger to the writer
		Logger: l,
	})

	kl.kafkaWriter = kWriter
	fmt.Println(kl)
}

// sends message to kafka
func (kl *KafkaLogger) ProduceMessage(k, val string) error {
	w := kl.kafkaWriter
	ctx := kl.kafkaCtx

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

// func SetupKafka() {
// 	KafkaCtx = context.Background()

// }
