package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Auth struct {

	Id			primitive.ObjectID		`json:"id" bson:"_id,omitempty"`
	Wallet		string					`json:"wallet" validate:"required"`
	Message 	string					`json:"message" validate:"required"`
	Signature	string					`json:"signature" validate:"required"`

}