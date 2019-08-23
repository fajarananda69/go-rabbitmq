# go-rabbitmq
RabbitMQ speaks multiple protocols. This tutorial uses AMQP 0-9-1, which is an open, general-purpose protocol for messaging. There are a number of clients for RabbitMQ in many different languages. We'll use the Go amqp client in this tutorial.

## Install and SETUP GO
Download and configure your workspace with latest version of Go and correct environment path.
- [Last Version](https://golang.org/dl/)
- [Windows](http://www.wadewegner.com/2014/12/easy-go-programming-setup-for-windows/)
- [Linux](http://www.tecmint.com/install-go-in-linux/)

## install and SETUP RabbitMQ
- install rabbitmq and rabbitmq management [klik here](https://www.rabbitmq.com/download.html)

## Install Driver
Install dep [klik here](https://golang.github.io/dep/docs/installation.html) and Create new dep
```
dep init 
```
Install driver rabbitmq
```
dep ensure -add github.com/streadway/amqp
```

## Import
```
import (
	"log"

	"github.com/streadway/amqp"
)
``` 
## Connection
```
func Connections() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://user:pass@localhost:5672/")
	return conn, err
}
```

## Publisher & Consumer
- Brokers and Their Role

Messaging brokers receive messages from publishers (applications that publish them, also known as producers) and route them to consumers (applications that process them).

Since it is a network protocol, the publishers, consumers and the broker can all reside on different machines

Let's quickly go over what we covered in the previous tutorials:

- A producer is a user application that sends messages.
- A queue is a buffer that stores messages.
- A consumer is a user application that receives messages.

### Set Exchange
```
err = ch.ExchangeDeclare(
		"jsa.rabbitmq", // name
		"topic",        // type
		true,           // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	)
```

### Set Router
```
err = ch.QueueBind(
		"jsa.queues",   // queue name
		"log.sys",      // routing key
		"jsa.rabbitmq", // exchange
		false,
		nil,
	)
```

### Publish
```
err = ch.Publish(
		"jsa.rabbitmq", // exchange
		"log.sys",      // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
```

### Set Queue
```
q, err := ch.QueueDeclare(
		"jsa.queues", // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
```

### Consume
```
msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
```

## Run
- run consumer
```
go run consumer.go
```
- open another terminal to run producer file `main.go`
```
go run main.go <args>
```
`<args>` => message for sent to rabbitmq queue 