package producer

import (
	rabbit "github.com/aminyasser/chat-api/golang-service/clients/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Produce(topic string, body []byte) error {
	channel := rabbit.GetRabbitMQPublishChannel()
	message := amqp.Publishing{
		ContentType: "json/application",
		Body:        body,
	}
	err := channel.Publish("", topic, true, false, message)
	return err
}