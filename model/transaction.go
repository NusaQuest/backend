package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transaction struct {

	Id				primitive.ObjectID		`json:"id" bson:"_id,omitempty"`
	Wallet			string					`json:"wallet"`
	Title			string					`json:"title"`
	Detail			string					`json:"detail"`
	TxHash			string					`json:"txHash"`
	TxTimestamp		uint64					`json:"txTimestamp"`

}