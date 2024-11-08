package feedback

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Feedback struct {
	ID           bson.ObjectID `bson:"_id" json:"id"`
	ArticleId    string        `bson:"articleId" json:"articleId" validate:"required,min=1,max=100"`
	CustomerName string        `bson:"customerName" json:"customerName" validate:"required,min=1,max=100"`
	FeedbackInfo string        `bson:"feedbackInfo" json:"feedbackInfo" validate:"required,min=1,max=100"`
	Rating       int           `bson:"rating" json:"rating" validate:"required,min=1,max=5"`
	CreationDate time.Time     `bson:"creationDate" json:"creationDate"`
}

type CreateFeedbackDto struct {
	FeedbackInfo string `bson:"feedbackInfo" json:"feedbackInfo" validate:"required,min=1,max=100"`
	Rating       int    `bson:"rating" json:"rating" validate:"required,min=1,max=5"`
	ArticleId    string `bson:"articleId" json:"articleId" validate:"required,min=1,max=100"`
}

type GetFeedbackDto struct {
	Feedbacks []struct {
		CustomerName string `json:"customerName"`
		FeedbackInfo string `json:"comentario del usuario"`
		Rating       int    `json:"calificación"`
		CreationDate string `json:"creationDate"`
	} `json:"feedbacks"`
}
