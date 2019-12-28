package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Produce(brokerAddress string, kafkaTopic string, message string) {

	producer, errorIfProducerNotCreated := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": brokerAddress})

	if errorIfProducerNotCreated != nil {
		fmt.Println("Could not create Kafka Producer", errorIfProducerNotCreated)
	}

	deliveryChannel := make(chan kafka.Event)

	fmt.Println("message", message)

	err := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &kafkaTopic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, deliveryChannel)

	if err != nil {
		fmt.Println("Error in producing message to kafka", err)
	}

	e := <-deliveryChannel
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s to partition %d at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}

	close(deliveryChannel)
}
