package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func produce(brokerAddress string, kafkaTopic string, message string) {

	producer, errorIfProducerNotCreated := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": brokerAddress})

	if errorIfProducerNotCreated != nil {
		fmt.Println("Could not create Kafka Producer", errorIfProducerNotCreated)
	}

	deliveryChan := make(chan kafka.Event)

	err := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &kafkaTopic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
		Headers:        []kafka.Header{{Key: "Dummy Key", Value: []byte("Dummy Value")}},
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
