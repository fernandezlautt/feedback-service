package feedback

import (
	"fernandezlautt/feedback-service/rabbit"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func findFeedbacks() ([]Feedback, error) {
	return findAll()
}

func createFeedback(ctx *gin.Context, feedbackDto CreateFeedbackDto) error {
	_, err := rabbit.RpcArticleExist(ctx, feedbackDto.ArticleId)
	if err != nil {
		return err
	}

	err = insert(Feedback{
		ID:           bson.NewObjectID(),
		ArticleId:    feedbackDto.ArticleId,
		CustomerName: "Name",
		FeedbackInfo: feedbackDto.FeedbackInfo,
		Rating:       feedbackDto.Rating,
		CreationDate: time.Now(),
	})

	if err != nil {
		return err
	}

	return nil
}
