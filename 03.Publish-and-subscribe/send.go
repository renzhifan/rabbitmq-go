package main

import (
	"github.com/streadway/amqp"
	"log"
	"rabbitmq-test/utils"
)

func main() {
	// 连接rabbitmq
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/guest")
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 创建信道
	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 声明交换机
	err = ch.ExchangeDeclare(
		"exchangeTest1",   // 交换机名字
		"fanout", // 交换机类型，fanout发布订阅模式
		true,     // 是否持久化
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	utils.FailOnError(err, "Failed to declare an exchange")

	// 消息内容
	body := "Hello abc"
	// 推送消息
	err = ch.Publish(
		"exchangeTest1",     // exchange（交换机名字，跟前面声明对应）
		"", // 路由参数，fanout类型交换机，自动忽略路由参数，填了也没用。
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/plain", // 消息内容类型，这里是普通文本
			Body:        []byte(body),  // 消息内容
		})

	log.Printf("发送内容 %s", body)
}