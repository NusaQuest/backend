package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Destination struct {

	Id				primitive.ObjectID			`json:"id" bson:"_id,omitempty"`
	Name			string						`json:"name" validate:"required"`
	City			string						`json:"city" validate:"required"`
	Province		string						`json:"province" validate:"required"`
	StreetName		string						`json:"streetName" validate:"required"`
	Map				string						`json:"map" validate:"required"`
	Images			[]string					`json:"images" validate:"required"`
	
}