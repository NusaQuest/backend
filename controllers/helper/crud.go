package helper

import (
	"context"
	"time"

	"github.com/NusaQuest/backend.git/config"
	"github.com/NusaQuest/backend.git/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
// @dev Parses the request body, validates the struct, and inserts it.
// @param c - Fiber context to parse the request body.
// @param doc - The name of the MongoDB collection.
// @param obj - The object to insert.
// @return The result of the insert operation and an error if occurred.
func InsertData(c *fiber.Ctx, doc string, obj interface{}) (*mongo.InsertOneResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := c.BodyParser(obj)
	if err != nil {
		return nil, err
	}

	err = utils.GetValidator().Struct(obj)
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
// @dev Parses the request body and updates fields with $set operator.
// @param c - Fiber context to parse the request body.
// @param doc - The name of the MongoDB collection.
// @param id - The ObjectID of the document to update.
// @param obj - The object containing updated fields.
// @return The result of the update operation and an error if occurred.
func UpdateData(c *fiber.Ctx, doc string, id primitive.ObjectID, obj interface{}) (*mongo.UpdateResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := c.BodyParser(obj)
	if err != nil {
		return nil, err
	}

	collection := config.GetDatabase().Collection(doc)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": obj}

	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return res, nil

}
