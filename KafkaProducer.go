package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func produce() {

	producer, errorIfProducerNotCreated := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "127.0.0.1:9092"})

	if errorIfProducerNotCreated != nil {
		fmt.Println("Could not create Kafka Producer", errorIfProducerNotCreated)
	}

	topic := "test"

	deliveryChan := make(chan kafka.Event)

	value := "From Kafka Tail!!"

	err := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(value),
		Headers:        []kafka.Header{{Key: "myTestHeader", Value: []byte("header values are binary")}},
	}, deliveryChan)

	if err != nil {
		fmt.Println("Error in producing message to kafka", err)
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s to partition %d at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}

	close(deliveryChan)
}
