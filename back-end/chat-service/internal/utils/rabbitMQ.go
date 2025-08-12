package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func waitForRabbitMQ(url string, maxRetries int, delay time.Duration) (*amqp.Connection, error) {
	for i := 0; i < maxRetries; i++ {
		conn, err := amqp.Dial(url)
		if err == nil {
			return conn, nil
		}
		log.Printf("RabbitMQ not ready, retrying in %v... (%d/%d)", delay, i+1, maxRetries)
		time.Sleep(delay)
	}
	return nil, fmt.Errorf("failed to connect to RabbitMQ after %d retries", maxRetries)
}

func RabbitMQ() (*amqp.Connection, *amqp.Channel, error) {
	USERNAME := os.Getenv("RABBITMQ_USERNAME")
	PASSWORD := os.Getenv("RABBITMQ_PASSWORD")
	HOST := os.Getenv("RABBITMQ_HOST")
	PORT := os.Getenv("RABBITMQ_PORT")

	URL := fmt.Sprintf("amqp://%s:%s@%s:%s/", USERNAME, PASSWORD, HOST, PORT)

	conn, err := waitForRabbitMQ(URL, 0, 2*time.Second)

	if err != nil {
		return &amqp.Connection{}, &amqp.Channel{}, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return &amqp.Connection{}, &amqp.Channel{}, err
	}

	return conn, ch, nil
}

func PublishMessage(ch *amqp.Channel, nameQueue string, content string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	q, err := ch.QueueDeclare(
		nameQueue, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return fmt.Errorf("Failed to declare a queue")
	}

	// Serialize to []byte
	data, err := json.Marshal(content)
	if err != nil {
		log.Fatal(err)
	}

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType:  "text/plain",
			DeliveryMode: amqp.Persistent,
			Body:         data,
		})

	if err != nil {
		return fmt.Errorf("Failed to publish a message")
	}

	return nil
}

func SubscribeMessage(ch *amqp.Channel, receiverId string) error {
	q, err := ch.QueueDeclare(
		receiverId, // name
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		return fmt.Errorf("Failed to declare a queue")
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		return fmt.Errorf("Failed to subscribe a message")
	}

	go func() {
		// for _ := range msgs {
		// 	// var payload models.ReqSendNotify
		// 	// _ = json.Unmarshal(d.Body, &payload)
		// 	// SendToWebSocket(payload.ReceiverId.String(), payload.Messages)
		// }

		fmt.Println(msgs)
	}()

	return nil
}
