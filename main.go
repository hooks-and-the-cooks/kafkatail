package main

import (
	"fmt"
	"github.com/kafkatail/kafka"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := cli.App{
		Name:        "kafkatail",
		Usage:       "Kafka CLI",
		Version:     "0.1.0",
		Description: "Kafka Consumer and Producer CLI in golang",
		Copyright:   "Ujjawal Dixit",
		Commands: []*cli.Command{
			{
				Name:        "producer",
				Aliases:     []string{"p"},
				Usage:       "Run Kafka tail in producer mode",
				Description: "To Produce message to kafka",
				Action: func(context *cli.Context) error {
					brokerAddress := context.Args().Get(0)
					kafkaTopic := context.Args().Get(1)
					message := context.Args().Get(2)
					kafka.Produce(brokerAddress, kafkaTopic, message)
					return nil
				},
				Flags: nil,
			},
			{
				Name:        "consumer",
				Aliases:     []string{"c"},
				Usage:       "Run Kafka tail in consumer mode",
				Description: "To Consume message from kafka",
				Action: func(context *cli.Context) error {
					brokerAddress := context.Args().Get(0)
					kafkaTopic := context.Args().Get(1)
					kafka.Consume(brokerAddress, kafkaTopic, false)
					return nil
				},
				Flags: nil,
			},
			{
				Name:        "consumer-from-bytes",
				Aliases:     []string{"cfb"},
				Usage:       "Run Kafka tail in consumer mode and messages in bytes",
				Description: "To Consume message from kafka",
				Action: func(context *cli.Context) error {
					brokerAddress := context.Args().Get(0)
					kafkaTopic := context.Args().Get(1)
					kafka.Consume(brokerAddress, kafkaTopic, true)
					return nil
				},
				Flags: nil,
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		fmt.Println("Something wrong happened at startup!")
	}
}
