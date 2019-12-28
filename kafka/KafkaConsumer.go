package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"os/signal"
	"syscall"
)

func Consume(brokerAddress string, topic string) {

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	consumer, _ := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":     brokerAddress,
		"broker.address.family": "v4",
		"group.id":              "consumer-group-1",
		"session.timeout.ms":    6000,
		"auto.offset.reset":     "latest"})

	topics := []string{}

	topics = append(topics, topic)

	_ = consumer.SubscribeTopics(topics, nil)

	run := true

	for run == true {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := consumer.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				fmt.Printf("%% Message on %s:\n%s\n",
					e.TopicPartition, string(e.Value))
				if e.Headers != nil {
					fmt.Printf("%% Headers: %v\n", e.Headers)
				}
			case kafka.Error:
				_, _ = fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)
				if e.Code() == kafka.ErrAllBrokersDown {
					run = false
				}
			default:
				fmt.Println("Polling Continuosly, Please produce something")
			}
		}
	}

	fmt.Printf("Closing consumer\n")
	_ = consumer.Close()
}
