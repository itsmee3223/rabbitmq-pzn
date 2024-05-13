package main

import (
	"context"
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

// format dial
// protocol amqp
// username
// password
// host
// port
// vhost jika / maka default vhost
// example: amqp://guest:guest@localhost:5672/

func main() {
	connection, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	defer connection.Close()

	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	// buat consumer dari channel
	// param 1 untuk context
	// param 2 untuk queue string
	emailConsumer, err := channel.ConsumeWithContext(ctx, "email", "email-consumer", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	// get semua message dari queues
	for message := range emailConsumer {
		fmt.Println("Routing key:", message.RoutingKey)
		fmt.Println("Body:", string(message.Body))
	}

}
