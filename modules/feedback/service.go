package feedback

import (
	"fernandezlautt/feedback-service/lib"
	"fernandezlautt/feedback-service/rabbit"
	"fernandezlautt/feedback-service/security"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func findFeedbacks(articleId string, page int, size int) ([]Feedback, error) {
	findWhere := bson.M{"articleId": articleId, "status": "confirmed"}
	findOrder := bson.M{"creationDate": -1}
	pagination := lib.GetPagination(page, size)
	fmt.Println(pagination)
	return findAll(findWhere, findOrder, pagination)
}

func createFeedback(ctx *gin.Context, feedbackDto CreateFeedbackDto, user *security.User) error {
	if feedbackDto.ArticleId == "" {
		return lib.ArticleIdRequired
	}

	res, err := insert(Feedback{
		ID:           bson.NewObjectID(),
		ArticleId:    feedbackDto.ArticleId,
		CustomerName: user.Name,
		CustomerId:   user.ID,
		FeedbackInfo: feedbackDto.FeedbackInfo,
		Rating:       feedbackDto.Rating,
		CreationDate: time.Now(),
		Status:       "pending",
	})

	if err != nil {
		return err
	}

	feedbackId := res.InsertedID.(bson.ObjectID).Hex()

	err = rabbit.SendArticleExist(ctx, feedbackDto.ArticleId, feedbackId)

	if err != nil {
		return err
	}

	return nil
}

func ConfirmFeedback(feedbackId string) error {
	return update(feedbackId, bson.M{"status": "confirmed"})
}

func DisableFeedback(feedbackId string, reason string) error {
	return update(feedbackId, bson.M{"status": "disabled", "reason": reason})
}
