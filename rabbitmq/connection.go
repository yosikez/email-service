package rabbitmq

import (
	"github.com/yosikez/email-notification-services/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

func NewRabbitMQ() (*config.RabbitMQ, *config.RabbitMQConnection, error) {
	var err error
	var rmq config.RabbitMQConnection

	rmqCfg, err := config.LoadRabbitMQ()

	if err != nil {
		return nil, nil, err
	}

	rabbitmqUrl := rmqCfg.GetUrl()

	rmq.Connection, err = amqp.Dial(rabbitmqUrl)

	if err != nil {
		return nil, nil, err
	}

	rmq.Channel, err = rmq.Connection.Channel()

	if err != nil {
		return nil, nil, err
	}

	return rmqCfg, &rmq, nil
	
}