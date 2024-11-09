package feedback

import (
	"context"
	"fernandezlautt/feedback-service/db"
	"fernandezlautt/feedback-service/lib"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var collection *mongo.Collection

func dbCollection() *mongo.Collection {
	if collection == nil {
		database := db.Get()
		collection = database.Collection("feedback")
	}

	return collection
}

func findAll(where bson.M, order bson.M, pagination *lib.Pagination) ([]Feedback, error) {
	var feedbacks []Feedback

	optionsFind := options.Find().SetSort(order).SetSkip(pagination.Skip).SetLimit(pagination.Limit)
	cursor, err := dbCollection().Find(context.TODO(), where, optionsFind)

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
