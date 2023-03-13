package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yosikez/email-notification-services/database"
	"github.com/yosikez/email-notification-services/mail"
	"github.com/yosikez/email-notification-services/rabbitmq"
	"github.com/yosikez/email-notification-services/router"
)

func main() {
	err := database.Connect()
	if err != nil {
		panic(err)
	}

	rmqCfg, rmq, err := rabbitmq.NewRabbitMQ()
	if err != nil {
		log.Fatalf("failed to connect to rabbitmq : %v", err)
	}

	defer rmq.Connection.Close()
	defer rmq.Channel.Close()

	err = rmq.Channel.ExchangeDeclare(
		rmqCfg.ExchangeName,
		rmqCfg.ExchangeKind,
		true,
		false,
		false,
		false,
		nil,
	)

	go mail.StartConsume(rmq, rmqCfg)

	if err != nil {
		log.Fatalf("failed to declare exchange : %v", err)
	}

	r := gin.Default()
	
	router.RegisterRouter(r)

	err = r.Run(":8080")

	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}

	fmt.Println("server started on port 8000")

}
