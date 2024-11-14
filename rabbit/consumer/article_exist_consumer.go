package consumer

import (
	"encoding/json"
	"fernandezlautt/feedback-service/lib"
	"fernandezlautt/feedback-service/modules/feedback"
	"fernandezlautt/feedback-service/rabbit"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ArticleExistConsume() {
	conn, err := amqp.Dial(lib.GetEnv().RabbitURL)
	defer conn.Close()
	if err != nil {
		log.Panic(err)
		return
	}
	ch, err := conn.Channel()
	defer ch.Close()
	if err != nil {
		log.Panic(err)
		return
	}
	queue, err := ch.QueueDeclare(
		"catalog_article_exist_receive", // name
		false,                           // durable
		false,                           // delete when unused
		false,                           // exclusive
		false,                           // no-wait
		nil,                             // arguments
	)
	if err != nil {
		log.Panic(err)
		return
	}

	err = ch.ExchangeDeclare(
		"article_exist", // name
		"direct",        // type
		false,           // durable
		false,           // auto-deleted
		false,           // internal
		false,           // no-wait
		nil,             // arguments
	)

	if err != nil {
		log.Panic(err)
		return
	}

	err = ch.QueueBind(
		queue.Name,              // queue name
		"article_exist_receive", // routing key
		"article_exist",         // exchange
		false,
		nil)

	if err != nil {
		log.Panic(err)
		return
	}

	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)

	if err != nil {
		log.Panic(err)
		return
	}

	for d := range msgs {
		var res rabbit.ConsumeSendArticleExist
		err := json.Unmarshal(d.Body, &res)
		if err != nil {
			log.Panic(err)
			continue
		}
		fmt.Printf("Received a message to update feedback: %s\n", res.CorrelationId)
		if res.Message.Valid {
			feedback.ConfirmFeedback(res.CorrelationId)
		} else {
			feedback.DisableFeedback(res.CorrelationId, "Articulo inexistente")
		}
	}
}
