package service

import (
	"NextEraAbyss/FiberForge/pkg/email"
	"NextEraAbyss/FiberForge/pkg/logger"
	"NextEraAbyss/FiberForge/pkg/rabbitmq"
	"encoding/json"
)

const (
	EmailQueueName = "email_queue"
)

type EmailMessage struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type EmailService struct{}

func NewEmailService() *EmailService {
	// 声明队列
	_, err := rabbitmq.DeclareQueue(EmailQueueName)
	if err != nil {
		logger.Error("Failed to declare email queue", "error", err)
	}

	// 启动消费者
	err = rabbitmq.ConsumeMessages(EmailQueueName, handleEmailMessage)
	if err != nil {
		logger.Error("Failed to start email consumer", "error", err)
	}

	return &EmailService{}
}

// SendEmail 发送邮件
func (s *EmailService) SendEmail(to, subject, body string) error {
	msg := EmailMessage{
		To:      to,
		Subject: subject,
		Body:    body,
	}

	// 序列化消息
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		logger.Error("Failed to marshal email message", "error", err)
		return err
	}

	// 发送消息到队列
	err = rabbitmq.PublishMessage(EmailQueueName, msgBytes)
	if err != nil {
		logger.Error("Failed to publish email message", "error", err)
		return err
	}

	logger.Info("Email message queued",
		"to", to,
		"subject", subject,
	)
	return nil
}

// handleEmailMessage 处理邮件消息
func handleEmailMessage(body []byte) error {
	var msg EmailMessage
	if err := json.Unmarshal(body, &msg); err != nil {
		logger.Error("Failed to unmarshal email message", "error", err)
		return err
	}

	// 使用邮件工具包发送邮件
	err := email.SendEmail(msg.To, msg.Subject, msg.Body)
	if err != nil {
		logger.Error("Failed to send email",
			"to", msg.To,
			"subject", msg.Subject,
			"error", err,
		)
		return err
	}

	logger.Info("Email sent successfully",
		"to", msg.To,
		"subject", msg.Subject,
	)
	return nil
}
