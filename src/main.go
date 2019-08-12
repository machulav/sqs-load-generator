package main

import (
	"github.com/machulav/sqs-load-generator/src/sender"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	sqsURL         = kingpin.Arg("sqs-url", "SQS queue URL").Required().String()
	sqsMessage     = kingpin.Arg("sqs-message", "SQS message payload").Required().String()
	parallelism    = kingpin.Flag("parallelism", "Number of parallel SQS message senders").Short('p').Default("1").Int()
	messagesAmount = kingpin.Flag("messages-amount", "Amount of messages that should sent to the SQS queue by one sender").Short('m').Default("1").Int()
)

func main() {
	kingpin.Parse()

	s := sender.Config{
		SqsURL:         *sqsURL,
		SqsMessage:     *sqsMessage,
		Parallelism:    *parallelism,
		MessagesAmount: *messagesAmount,
	}

	s.Run()
}
