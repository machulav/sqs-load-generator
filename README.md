# sqs-load-generator

A tool that allows adding the required amount of a specified message into SQS queue.

It was created for SQS consumer load testing.

## Build

```
go build -o main ./src
```

## Run

```
usage: main [<flags>] <sqs-url> <sqs-message>

Flags:
      --help               Show context-sensitive help
  -p, --parallelism=1      Number of parallel SQS message senders
  -m, --messages-amount=1  Amount of messages that should be sent to the SQS queue by one sender

Args:
  <sqs-url>      SQS queue URL
  <sqs-message>  SQS message payload
```