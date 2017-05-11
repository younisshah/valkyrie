package vrabbit

import (
	"log"

	"github.com/streadway/amqp"
)

/**
*  Created by Galileo on 12/5/17.
*
*  RabbitMQ Client
 */

const _RABBIT_MQ_URL = "amqp://guest:guest@localhost:5672/"

type _RabbitMQ struct {
	url       string
	queueName string
	conn      *amqp.Connection
}

func NewRabbitMQConfig(queueName string) *_RabbitMQ {
	c := &_RabbitMQ{url: _RABBIT_MQ_URL}
	if len(queueName) == 0 {
		c.queueName = "valkyrie_queue"
	} else {
		c.queueName = queueName
	}
	connection, err := c.getConnection()
	failRabbitOnErr(err)

	c.conn = connection
	return c
}

func (r *_RabbitMQ) Produce(message string) error {
	channel, err := r.conn.Channel()
	failRabbitOnErr(err)

	q, err := r.declareQueue(r.queueName, channel)
	failRabbitOnErr(err)

	return r.publish(message, q, channel)
}

func (r *_RabbitMQ) Close() {
	r.conn.Close()
}

func (r *_RabbitMQ) getConnection() (*amqp.Connection, error) {
	return amqp.Dial(r.url)
}

func (r *_RabbitMQ) declareQueue(queueName string, channel *amqp.Channel) (amqp.Queue, error) {
	return channel.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil)
}

func (r *_RabbitMQ) publish(message string, queue amqp.Queue, channel *amqp.Channel) error {
	return channel.Publish("",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

func failRabbitOnErr(err error) {
	if err != nil {
		log.Fatal("[*] RabbitMQ error:", err)
	}
}
