package main

import (
	producer "go-rabbitMq/rabbitmq/producer"
	"os"
)

func main() {

	producer.Producer(os.Args[1])
	// consumer.Consumer()
}
