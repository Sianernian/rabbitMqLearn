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

	q, err := ch.QueueDeclare(
		"quan",
		false, //是否持久化存储（磁盘） 默认 存在内存中
		false, //
		false, //是否供一个消费者进行消费，是否进行消息共享
		false,
		nil)
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
