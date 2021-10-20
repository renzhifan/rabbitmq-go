package main

import (
	"github.com/streadway/amqp"
	"log"
	"rabbitmq-test/utils"
	"time"
)

func main() {
	// 连接RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/guest")
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 通过协程创建5个消费者
	for i := 0; i < 5; i++ {
		go func(number int) {
			// 创建一个rabbitmq信道, 每个消费者一个
			ch, err := conn.Channel()
			utils.FailOnError(err, "Failed to open a channel")
			defer ch.Close()

			// 声明需要操作的队列
			q, err := ch.QueueDeclare(
				"hello", // 队列名
				false,   // 是否需要持久化
				false,   // delete when unused
				false,   // exclusive
				false,   // no-wait
				nil,     // arguments
			)
			utils.FailOnError(err, "Failed to declare a queue")

			// 创建一个消费者
			msgs, err := ch.Consume(
				q.Name, // 需要操作的队列名
				"",     // 消费者唯一id，不填，则自动生成一个唯一值
				true,   // 自动提交消息（即自动确认消息已经处理完成）
				false,  // exclusive
				false,  // no-local
				false,  // no-wait
				nil,    // args
			)
			utils.FailOnError(err, "Failed to register a consumer")

			// 循环处理消息
			for d := range msgs {
				log.Printf("[消费者编号=%d] 收到消息: %s", number, d.Body)
				// 模拟业务处理，休眠1秒
				time.Sleep(time.Second)
			}
		}(i)
	}

	// 挂起主协程，避免程序退出
	forever := make(chan bool)
	<-forever
}