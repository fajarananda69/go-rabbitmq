package producer

import (
	dao "go-rabbitMq/rabbitmq/dao"
	"log"

	"github.com/streadway/amqp"
)

func Producer(args string) {
	conn, err := dao.Connections()
	dao.FailOnError(err, "Connection Failed")
	// defer conn.Close()

	ch, err := conn.Channel()
	dao.FailOnError(err, "Channel Failed")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"jsa.rabbitmq", // name
		"topic",        // type
		true,           // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	)
	dao.FailOnError(err, "Failed to declare an exchange")

	err = ch.QueueBind(
		"jsa.queues",   // queue name
		"log.sys",      // routing key
		"jsa.rabbitmq", // exchange
		false,
		nil,
	)

	body := args
	err = ch.Publish(
		"jsa.rabbitmq", // exchange
		"log.sys",      // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	dao.FailOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)
}
