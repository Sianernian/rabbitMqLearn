package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"rabbitMB/utils"
)

func main() {
	conn := utils.RabbitMQConn()
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("produce1 conn.Channel err:%+v", err)
	}

	q, err := ch.QueueDeclare("an", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("ch.QueueDeclare err:%+v", err)
	}
	for i := 0; i < 100; i++ {
		messag := fmt.Sprintf("{\"order_id\":%d}", i)
		fmt.Println("body:", messag)
		err = ch.Publish("", q.Name, false, false, amqp.Publishing{
			ContentType: "text:plain",
			Body:        []byte(messag),
		})
		if err != nil {
			log.Fatalf("ch.publish err:%+v", err)
		}
	}
}
