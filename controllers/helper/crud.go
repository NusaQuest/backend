package helper

import (
	"context"
	"time"

	"github.com/NusaQuest/backend.git/config"
	"github.com/NusaQuest/backend.git/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// @notice Retrieve documents from a MongoDB collection based on a filter.
// @dev It deserializes the found documents into the provided obj interface.
// @param filter - BSON filter for querying documents.
// @param doc - The name of the MongoDB collection.
// @param obj - The target object to deserialize the results into.
// @return The deserialized object and an error if occurred.
func RetrieveData(filter bson.M, doc string, obj interface{}) (interface{}, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := config.GetDatabase().Collection(doc)
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, obj)
	if err != nil {
		return nil, err
	}

	return obj, nil

}

// @notice Insert a new document into a MongoDB collection.
// @dev Validates the struct, and inserts it.
// @param c - Fiber context to parse the request body.
// @param doc - The name of the MongoDB collection.
// @param obj - The object to insert.
// @return The result of the insert operation and an error if occurred.
func InsertData(doc string, obj interface{}) (*mongo.InsertOneResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := utils.GetValidator().Struct(obj)
	if err != nil {
		return nil, err
	}

	collection := config.GetDatabase().Collection(doc)
	res, err := collection.InsertOne(ctx, obj)
	if err != nil {
		return nil, err
	}

	return res, nil

}

// @notice Update an existing document in a MongoDB collection by ID.
// @dev Updates fields with $set operator.
// @param c - Fiber context to parse the request body.
// @param doc - The name of the MongoDB collection.
// @param id - The ObjectID of the document to update.
// @param obj - The object containing updated fields.
// @return The result of the update operation and an error if occurred.
func UpdateData(doc string, key string, value interface{}, obj interface{}) (*mongo.UpdateResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := config.GetDatabase().Collection(doc)
	filter := bson.M{key: value}
	update := bson.M{"$set": obj}

	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return res, nil

}

// @notice Delete a document from a MongoDB collection.
// @param doc The name of the MongoDB collection.
// @param key The field name to filter the document.
// @param value The value of the field to match the document.
// @return The result of the delete operation, or an error if it fails.
func DeleteData(doc string, key string, value interface{}) (*mongo.DeleteResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := config.GetDatabase().Collection(doc)
	filter := bson.M{key: value}

	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return res, nil

}