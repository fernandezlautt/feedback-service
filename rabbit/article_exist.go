package rabbit

import (
	"encoding/json"
	"fernandezlautt/feedback-service/lib"
	"log"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

func SendArticleExist(ctx *gin.Context, articleID string, feedbackId string) error {
	conn, err := amqp.Dial(lib.GetEnv().RabbitURL)
	defer conn.Close()
	if err != nil {
		log.Panic(err)
		return err
	}
	ch, err := conn.Channel()
	defer ch.Close()
	if err != nil {
		log.Panic(err)
		return err
	}

	toSend, err := json.Marshal(ArticleExistReq{
		CorrelationId: feedbackId,
		Exchange:      "article_exist",
		RoutingKey:    "article_exist_receive",
		Message: &ArticleExistMessage{
			ArticleId:   articleID,
			ReferenceId: articleID,
		}})

	if err != nil {
		log.Panic(err)
		return err
	}

	err = ch.PublishWithContext(ctx,
		"article_exist", // exchange
		"article_exist", // routing key
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: feedbackId,
			ReplyTo:       "article_exist_receive",
			Body:          toSend,
		})

	if err != nil {
		log.Panic(err)
		return err
	}

	return nil
}
