package vrabbit

import (
	"log"

	"fmt"

	"github.com/streadway/amqp"
)

/**
*  Created by Galileo on 12/5/17.
*
*  RabbitMQ Client
 */

type RabbitMQ struct {
	conn *amqp.Connection
}

func (r *RabbitMQ) Connect(url string) error {
	if connection, err := amqp.Dial(url); err != nil {
		return err
	} else {
		r.conn = connection
	}
	return nil
}

func (r *RabbitMQ) Produce(message interface{}, queueName string) error {
	channel, err := r.conn.Channel()
	failRabbitOnErr(err)

	q, err := r.declareQueue(queueName, channel)
	failRabbitOnErr(err)

	return r.publish(message, q, channel)
}

func (r *RabbitMQ) Consume(queueName string, callback func(data interface{}) error) {
	channel, err := r.conn.Channel()
	failRabbitOnErr(err)

	q, err := r.declareQueue(queueName, channel)
	failRabbitOnErr(err)

	failRabbitOnErr(r.channelQOS(channel))

	messages, err := channel.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	eternity := make(chan struct{})

	go func() {
		for m := range messages {
			body := string([]byte(m.Body))
			fmt.Println("[+] Consumed:")
			callback(body)
			m.Ack(false)
		}
	}()
	fmt.Println("[+] Consumer ready")
	<-eternity
}

func (r *RabbitMQ) Close() {
	r.conn.Close()
}

//***Helper methods****
func (r *RabbitMQ) channelQOS(channel *amqp.Channel) error {
	return channel.Qos(
		1,
		0,
		false,
	)
}

func (r *RabbitMQ) declareQueue(queueName string, channel *amqp.Channel) (amqp.Queue, error) {
	return channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
}

func (r *RabbitMQ) publish(message interface{}, queue amqp.Queue, channel *amqp.Channel) error {
	return channel.Publish("",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(message.(string)),
		})
}

func failRabbitOnErr(err error) {
	if err != nil {
		log.Fatal("[*] RabbitMQ error:", err)
	}
}
