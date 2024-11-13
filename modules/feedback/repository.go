package feedback

import (
	"context"
	"fernandezlautt/feedback-service/db"
	"fernandezlautt/feedback-service/lib"
	"fmt"

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

func insert(feedback Feedback) (*mongo.InsertOneResult, error) {
	result, err := dbCollection().InsertOne(context.TODO(), feedback)
	return result, err
}

// unused
func delete(feedbackId string) (*mongo.DeleteResult, error) {
	oid, err := bson.ObjectIDFromHex(feedbackId)
	if err != nil {
		return nil, err
	}
	where := bson.M{"_id": oid}
	return dbCollection().DeleteOne(context.TODO(), where)
}

func update(feedbackId string, update bson.M) error {
	oid, err := bson.ObjectIDFromHex(feedbackId)
	if err != nil {
		return err
	}
	where := bson.M{"_id": oid}
	updateQuery := bson.M{"$set": update}
	_, err = dbCollection().UpdateOne(context.TODO(), where, updateQuery)
	fmt.Println("congratsssss")
	return err
}
