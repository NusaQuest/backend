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

func RetrieveData(filter bson.M, doc string, obj interface{}) (interface{}, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := config.GetDatabase().Collection(doc)
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &obj)
	if err != nil {
		return nil, err
	}

	return obj, nil

}

func InsertData(c *fiber.Ctx, doc string, obj interface{}) (*mongo.InsertOneResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := c.BodyParser(&obj)
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

func UpdateData(c *fiber.Ctx, id primitive.ObjectID, obj interface{}) (*mongo.UpdateResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := c.BodyParser(&obj)
	if err != nil {
		return nil, err
	}

	collection := config.GetDatabase().Collection("proposals")
	filter := bson.M{"_id": id}
	update := bson.M{"$set": obj}

	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return res, nil

}
