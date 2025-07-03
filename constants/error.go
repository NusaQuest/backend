package constants

type Error string

const (
	DatabaseNotInitialized   Error = "MongoDB hasn’t been initialized or connected properly. Please check the database setup."
	ErrorWhilePingToDatabase Error = "Unable to reach the MongoDB database. Please ensure it is running and accessible."
)
