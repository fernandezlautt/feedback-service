package feedback

import (
	"fernandezlautt/feedback-service/lib"
	"fernandezlautt/feedback-service/rabbit"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func findFeedbacks(articleId string, page int, size int) ([]Feedback, error) {
	findWhere := bson.M{"articleId": articleId}
	findOrder := bson.M{"creationDate": -1}
	pagination := lib.GetPagination(page, size)
	fmt.Println(pagination)
	return findAll(findWhere, findOrder, pagination)
}

func createFeedback(ctx *gin.Context, feedbackDto CreateFeedbackDto, userName string, articleId string) error {
	_, err := rabbit.RpcArticleExist(ctx, articleId)
	if err != nil {
		return err
	}

	err = insert(Feedback{
		ID:           bson.NewObjectID(),
		ArticleId:    articleId,
		CustomerName: userName,
		FeedbackInfo: feedbackDto.FeedbackInfo,
		Rating:       feedbackDto.Rating,
		CreationDate: time.Now(),
	})

	if err != nil {
		return err
	}

	return nil
}
