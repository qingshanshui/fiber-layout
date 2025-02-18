package rabbitmq

import (
	"NextEraAbyss/FiberForge/pkg/config"
	"NextEraAbyss/FiberForge/pkg/logger"
	"context"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	conn    *amqp.Connection
	channel *amqp.Channel
)

// Init 初始化 RabbitMQ 连接
func Init() {
	if !config.Conf.RabbitMQ.Enable {
		logger.Info("RabbitMQ is disabled")
		return
	}

	var err error
	// 构建连接URL
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/",
		config.Conf.RabbitMQ.Username,
		config.Conf.RabbitMQ.Password,
		config.Conf.RabbitMQ.Host,
		config.Conf.RabbitMQ.Port,
	)

	// 建立连接
	conn, err = amqp.Dial(url)
	if err != nil {
		logger.Fatal("Failed to connect to RabbitMQ", "error", err)
	}

	// 创建通道
	channel, err = conn.Channel()
	if err != nil {
		logger.Fatal("Failed to open channel", "error", err)
	}

	logger.Info("Successfully connected to RabbitMQ",
		"host", config.Conf.RabbitMQ.Host,
		"port", config.Conf.RabbitMQ.Port,
	)
}

// DeclareQueue 声明队列
func DeclareQueue(name string) (amqp.Queue, error) {
	return channel.QueueDeclare(
		name,  // 队列名称
		true,  // 持久化
		false, // 自动删除
		false, // 排他性
		false, // 不等待
		nil,   // 参数
	)
}

// PublishMessage 发布消息
func PublishMessage(queueName string, body []byte) error {
	if channel == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return channel.PublishWithContext(ctx,
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
}

// ConsumeMessages 消费消息
func ConsumeMessages(queueName string, handler func([]byte) error) error {
	if channel == nil {
		return nil
	}

	msgs, err := channel.Consume(
		queueName, // 队列名称
		"",        // 消费者
		true,      // 自动确认
		false,     // 排他性
		false,     // 不等待
		false,     // 不阻塞
		nil,       // 参数
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			if err := handler(msg.Body); err != nil {
				logger.Error("Failed to process message",
					"queue", queueName,
					"error", err,
				)
			}
		}
	}()

	return nil
}

// Close 关闭连接
func Close() {
	if channel != nil {
		channel.Close()
	}
	if conn != nil {
		conn.Close()
	}
}
