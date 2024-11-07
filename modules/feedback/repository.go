package feedback

import (
	"context"
	"fernandezlautt/feedback-service/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var collection *mongo.Collection

func dbCollection() *mongo.Collection {

	if collection == nil {
		database := db.Get()
		collection = database.Collection("feedback")
	}

	return collection
}

func findAll() ([]Feedback, error) {
	var feedbacks []Feedback

	cursor, err := dbCollection().Find(context.TODO(), bson.M{})

	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &feedbacks); err != nil {
		return nil, err
	}

	return feedbacks, nil
}

func insert(feedback Feedback) error {
	_, err := dbCollection().InsertOne(context.TODO(), feedback)

	return err
}
