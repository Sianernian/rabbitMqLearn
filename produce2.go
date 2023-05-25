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
		log.Fatalf("conn.channel err:%+v", err)
	}

	que, err := ch.QueueDeclare("quan", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("channel.QueueDeclare err:%+v", err)
	}

	for i := 0; i < 100; i++ {
		message := fmt.Sprintf("{\"task_id\":%d}", i)
		fmt.Println("message:", message)
		err = ch.Publish("", que.Name, false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
		if err != nil {
			log.Fatalf("ch.Publish err:%v", err)
		}
	}
}
