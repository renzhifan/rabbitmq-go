package main

// 导入包
import (
	"github.com/streadway/amqp"
	"log"
)

// 错误处理
func failOnError11(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// 连接RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/guest")
	failOnError11(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 创建信道
	ch, err := conn.Channel()
	failOnError11(err, "Failed to open a channel")
	defer ch.Close()

	// 声明要操作的队列
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError11(err, "Failed to declare a queue")

	// 要发送的消息内容
	body := "Hello World!"

	// 发送消息
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError11(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}
