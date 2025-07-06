package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Identity struct {
	Id           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Wallet       string             `json:"wallet" validate:"required"`
	Hash         string             `json:"hash" validate:"required"`
	RegisteredAt int64              `json:"registeredAt" validate:"required"`
}
