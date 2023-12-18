package rabbitmq

import (
     "log"
	amqp "github.com/rabbitmq/amqp091-go"
)

var publishChannel *amqp.Channel
var consumeChannel *amqp.Channel


func GetRabbitMQPublishChannel() *amqp.Channel {
	if publishChannel == nil {
		initRabbitMQ()
	}
	return publishChannel
}
func GetRabbitMQConsumeChannel() *amqp.Channel {
	if consumeChannel == nil {
		initRabbitMQ()
	}
	return consumeChannel
}


func logOnError(err error, msg string) {
	if err != nil {
	  log.Printf("%s: %s", msg, err)
	}
}


func initRabbitMQ() {
	var err error
	rabbitConnection, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		logOnError(err, "Failed to connect to RabbitMQ")
	}
	publishChannel, err = rabbitConnection.Channel()
	if err != nil {
		logOnError(err, "Failed to open a channel")
	}
	_, err = publishChannel.QueueDeclare("chats_queue", true, false, false, false, nil)
	if err != nil {
		logOnError(err, "Failed to declare a queue")
	}
	// _, err = publishChannel.QueueDeclare(models.MessagesMQTopic, true, false, false, false, nil)
	// if err != nil {
	// 	logOnError(err, "Failed to declare a queue")
	// }
	consumeChannel, err = rabbitConnection.Channel()
	if err != nil {
		logOnError(err, "Failed to open a channel")
	}

	log.Println("Queue declared successfully.........")
}


