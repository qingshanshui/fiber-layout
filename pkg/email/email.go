package email

import (
	"fmt"
	"net/smtp"
	"wat.ink/layout/fiber/pkg/config"
	"wat.ink/layout/fiber/pkg/logger"
)

type EmailConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

var emailConfig EmailConfig

func Init() {
	emailConfig = EmailConfig{
		Host:     config.Conf.Email.Host,
		Port:     config.Conf.Email.Port,
		Username: config.Conf.Email.Username,
		Password: config.Conf.Email.Password,
		From:     config.Conf.Email.From,
	}
}

// SendEmail 发送邮件
func SendEmail(to, subject, body string) error {
	auth := smtp.PlainAuth("", emailConfig.Username, emailConfig.Password, emailConfig.Host)

	// 构建邮件内容
	msg := []byte(fmt.Sprintf("To: %s\r\n"+
		"From: %s\r\n"+
		"Subject: %s\r\n"+
		"Content-Type: text/plain; charset=UTF-8\r\n"+
		"\r\n"+
		"%s", to, emailConfig.From, subject, body))

	addr := fmt.Sprintf("%s:%d", emailConfig.Host, emailConfig.Port)
	if err := smtp.SendMail(addr, auth, emailConfig.From, []string{to}, msg); err != nil {
		logger.Error("Failed to send email",
			"to", to,
			"subject", subject,
			"error", err,
		)
		return err
	}

	logger.Info("Email sent successfully",
		"to", to,
		"subject", subject,
	)
	return nil
} 