package feedback

import "time"

type Feedback struct {
	ID           string    `bson:"_id" json:"id"`
	ArticleId    string    `bson:"articleId" json:"articleId" validate:"required,min=1,max=100"`
	CustomerName string    `bson:"customerName" json:"customerName" validate:"required,min=1,max=100"`
	FeedbackInfo string    `bson:"feedbackInfo" json:"feedbackInfo" validate:"required,min=1,max=100"`
	Rating       int       `bson:"rating" json:"rating" validate:"required,min=1,max=5"`
	CreationDate time.Time `bson:"creationDate" json:"creationDate"`
}

type CreateFeedbackDto struct {
	CustomerId   string `bson:"customerId" json:"customerId" validate:"required,min=1,max=100"`
	FeedbackInfo string `bson:"feedbackInfo" json:"feedbackInfo" validate:"required,min=1,max=100"`
	Rating       int    `bson:"rating" json:"rating" validate:"required,min=1,max=5"`
}

type GetFeedbackDto struct {
	Feedbacks []struct {
		CustomerName string `json:"customerName"`
		FeedbackInfo string `json:"comentario del usuario"`
		Rating       int    `json:"calificación"`
		CreationDate string `json:"creationDate"`
	} `json:"feedbacks"`
}
