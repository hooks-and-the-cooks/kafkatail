package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:  "b",
			Usage: "broker address",
			Value: "127.0.0.1:9092",
		}}

	app := cli.App{
		Name:        "kafkatail",
		Usage:       "Kafka CLI",
		Version:     "0.1.0",
		Description: "Kafka Consumer and Producer CLI in golang",
		Copyright:   "Ujjawal Dixit",
		Flags:       flags,
		Action: func(context *cli.Context) error {
			fmt.Println("Hello from cli!!")
			return nil
		}}

	err := app.Run(os.Args)

	if err != nil {
		fmt.Println("Something wrong happened at startup!")
	}
}
