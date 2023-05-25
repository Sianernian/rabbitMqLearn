package main

import (
	"log"
	"rabbitMB/utils"
	"time"
)

func main() {
	conn := utils.RabbitMQConn()
	defer conn.Close()

	forever := make(chan bool)

	for i := 0; i < 10; i++ {
		go func(num int) {
			// 获取通道  所有操作通过通道控制
			ch, err := conn.Channel()
			if err != nil {
				log.Fatalf("获取通道失败 conn.channel err:%+v", err)
			}
			defer ch.Close()
			// 队列声明
			q, err := ch.QueueDeclare(
				"an",
				false,
				false,
				false,
				false,
				nil,
			)
			if err != nil {
				log.Fatalf("队列声明失败 ch.QueueDeclare err:%+v", err)
			}
			// 设置每次从消息队列获取的数量
			err = ch.Qos(
				10,
				0,
				false, // 全局设置 为true时 全部channel适用  false时 这条通道使用
			)
			if err != nil {
				log.Fatalf("设置消息队列数量失败 ch.Qos err:%+v", err)
			}
			//消费者接受消息 msg为只读channel
			msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
			if err != nil {
				log.Fatalf("接受消息失败 ch.Consumer err；%+v", err)
			}
			for msg := range msgs {
				log.Printf("协程 %d  接受的消息：%s", num, msg.Body)
				time.Sleep(1 * time.Second)
			}
		}(i)
	}
	<-forever
}
