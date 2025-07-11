package config

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// / @notice Global MongoDB client instance
// @dev This client will be reused across different database operations
var Client *mongo.Client

// GetDatabase returns a reference to the configured MongoDB database.
// @notice Loads the DB_NAME from the .env file and returns the database instance.
// @dev This assumes ConnectDatabase has already been called to initialize the `Client`.
// @return *mongo.Database - The active database connection
func GetDatabase() *mongo.Database {

	DB_NAME := os.Getenv("DB_NAME")

	return Client.Database(DB_NAME)

}

// ConnectDatabase establishes a MongoDB connection and stores it in the global Client.
// @notice Loads MONGO_URI from .env and attempts to connect to MongoDB.
// @dev Should be called at application startup to initialize the database client.
// @return void - Logs a fatal error and exits if the connection fails
func ConnectDatabase() {

	MONGO_URI := os.Getenv("MONGO_URI")
	clientOption := options.Client().ApplyURI(MONGO_URI)

	client, err := mongo.Connect(context.Background(), clientOption)
	if err != nil {
		log.Fatalf("Error while connect to MongoDB!")
	}

	Client = client

	log.Println("Successfully connected to MongoDB!")

}

// DisconnectDatabase cleanly closes the MongoDB connection.
// @notice Should be called when the application is shutting down.
// @dev Ensures that resources are properly cleaned up.
// @return void - Logs a fatal error and exits if disconnection fails
func DisconnectDatabase() {

	err := Client.Disconnect(context.Background())
	if err != nil {
		log.Fatalf("Error while disconnect from MongoDB!")
	}

	log.Println("Successfully disconnected from MongoDB!")

}
