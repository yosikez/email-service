package mail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"strings"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/yosikez/email-notification-services/config"
	"github.com/yosikez/email-notification-services/database"
	"github.com/yosikez/email-notification-services/helper"
	"github.com/yosikez/email-notification-services/model"
	"gopkg.in/gomail.v2"
)

func dataSMTP() *config.SMTPConfig {
	smtpConfig, err := config.LoadSMTP()

	if err != nil {
		log.Fatalf("failed get data smtpconfig : %v", err)
	}

	return smtpConfig
}

func sendNotificationEmail(d *amqp.Delivery, action string) {

	SMTP := dataSMTP()

	var data helper.Message

	err := json.Unmarshal(d.Body, &data)

	if err != nil {
		log.Fatalf("failed to unmarshal : %v", err)
	}

	templateFile := fmt.Sprintf("template/%s.html", action)

	result, err := parseTemplate(templateFile, data)
	if err != nil {
		log.Fatalf("failed to parse template : %v", err)
	}

	title := strings.Title(action)
	subject := fmt.Sprintf("%s Todo Notification", title)

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", SMTP.Sender)
	mailer.SetHeader("To", data.UserEmail)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", result)

	dialer := gomail.NewDialer(
		SMTP.Host,
		SMTP.Port,
		SMTP.Username,
		SMTP.Password,
	)

	err = dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatalf("failed to send email : %v", err)
	}

	mail := &model.Mail{
		Action:   action,
		Receiver: data.UserEmail,
	}

	if err := database.DB.Create(&mail).Error; err != nil {
		log.Fatalf("failed to create record mail in database: %v", err)
	}

	log.Println("Email send successfully")

}

func parseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	if err := t.Execute(buf, data); err != nil {
		fmt.Println(err)
		return "", err
	}

	return buf.String(), nil
}

func StartConsume(rmq *config.RabbitMQConnection, rmqCfg *config.RabbitMQ) {
	go rmqDeclare("todo_create_queue", "create", rmq, rmqCfg)
	go rmqDeclare("todo_update_queue", "update", rmq, rmqCfg)
	go rmqDeclare("todo_done_queue", "done", rmq, rmqCfg)
	go rmqDeclare("todo_delete_queue", "delete", rmq, rmqCfg)
}

func rmqDeclare(queueName, action string, rmq *config.RabbitMQConnection, rmqCfg *config.RabbitMQ) {
	q, err := rmq.Channel.QueueDeclare(queueName, false, false, false, false, nil)

	if err != nil {
		log.Fatalf("failed to declare queue : %v", err)
	}

	err = rmq.Channel.QueueBind(q.Name, queueName, rmqCfg.ExchangeName, false, nil)

	if err != nil {
		log.Fatalf("failed to bind queue : %v", err)
	}

	msgs, err := rmq.Channel.Consume(queueName, "", true, false, false, false, nil)

	if err != nil {
		log.Fatalf("failed to register a consumer : %v", err)
	}

	for d := range msgs {
		sendNotificationEmail(&d, action)
	}
}