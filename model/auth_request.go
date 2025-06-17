package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type AuthRequest struct {

	Id			primitive.ObjectID		`json:"id" bson:"_id,omitempty"`
	Wallet		string					`json:"wallet"`
	Message 	string					`json:"message"`
	Signature	string					`json:"signature"`

}