package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transaction struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Wallet      string             `json:"wallet" validate:"required"`
	Title       string             `json:"title" validate:"required"`
	Detail      string             `json:"detail" validate:"required"`
	TxHash      string             `json:"txHash" validate:"required"`
	TxTimestamp uint64             `json:"txTimestamp" validate:"required"`
}
