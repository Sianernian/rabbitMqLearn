package utils

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func RabbitMQConn() (conn *amqp.Connection) {
	user := "an"
	pwd := "123456"
	host := "192.168.111.99"
	port := "5672"
	url := "amqp://" + user + ":" + pwd + "@" + host + ":" + port + "/"

	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("amqp.dial err:+%v", err)
	}
	return conn
}
