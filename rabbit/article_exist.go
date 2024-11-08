package rabbit

import (
	"encoding/json"
	"fernandezlautt/feedback-service/lib"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	uuid "github.com/satori/go.uuid"
)

func RpcArticleExist(ctx *gin.Context, articleID string) (*ConsumeArticleExistMessage, error) {
	conn, err := amqp.Dial(lib.GetEnv().RabbitURL)
	defer conn.Close()
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	ch, err := conn.Channel()
	defer ch.Close()
	if err != nil {
		log.Panic(err)
		return nil, err
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
		return nil, err
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
		return nil, err
	}

	err = ch.QueueBind(
		queue.Name,              // queue name
		"article_exist_receive", // routing key
		"article_exist",         // exchange
		false,
		nil)
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	CorrelationId := GetCorrelationId(ctx)

	toSend, err := json.Marshal(ArticleExistReq{
		CorrelationId: CorrelationId,
		Exchange:      "article_exist",
		RoutingKey:    "article_exist_receive",
		Message: &ArticleExistMessage{
			ArticleId:   articleID,
			ReferenceId: articleID,
		}})

	if err != nil {
		log.Panic(err)
		return nil, err
	}

	fmt.Println("Sending article exist message", CorrelationId)
	fmt.Println(string(toSend))
	err = ch.PublishWithContext(ctx,
		"article_exist", // exchange
		"article_exist", // routing key
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: CorrelationId,
			ReplyTo:       "article_exist_receive",
			Body:          toSend,
		})

	if err != nil {
		log.Panic(err)
		return nil, err
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
		return nil, err
	}

	for d := range msgs {
		var res ConsumeSendArticleExist
		err := json.Unmarshal(d.Body, &res)
		if err != nil {
			log.Panic(err)
			return nil, err
		}
		fmt.Println("sim")
		fmt.Println(res.CorrelationId)
		fmt.Println(CorrelationId)
		fmt.Println(res.Message)
		if res.CorrelationId != CorrelationId {
			continue
		}
		return &res.Message, nil
	}

	return nil, nil
}

func GetCorrelationId(c *gin.Context) string {
	value := c.GetHeader("correlation_id")

	if len(value) == 0 {
		value = uuid.NewV4().String()
	}

	return value
}
