package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/kafkatail/translator"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func Consume(brokerAddress string, topic string, isMessageInBytes bool) {

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	consumerGroupName := "consumer-group-2"
	consumer, _ := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":     brokerAddress,
		"broker.address.family": "v4",
		"group.id":              consumerGroupName,
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
			err := exec.Command("kafka-consumer-groups", "--bootstrap-server", brokerAddress,
				"--delete", "--group", consumerGroupName).Run()
			if err != nil {
				fmt.Println("kafka is not responding!!")
			}
			run = false
		default:
			ev := consumer.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				if isMessageInBytes {
					fmt.Printf("%% Message on %s:\n%s\n", e.TopicPartition, translator.Translate(e.Value))
				} else {
					fmt.Printf("%% Message on %s:\n%s\n", e.TopicPartition, string(e.Value))
				}
			case kafka.Error:
				_, _ = fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)
				if e.Code() == kafka.ErrAllBrokersDown {
					run = false
				}
			default:
				fmt.Println("Polling Continuously, Please produce something")
			}
		}
	}

	fmt.Printf("Closing consumer\n")
	_ = consumer.Close()
}
