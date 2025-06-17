package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Destination struct {

	Id				primitive.ObjectID			`json:"id" bson:"_id,omitempty"`
	Name			string						`json:"name"`
	City			string						`json:"city"`
	Province		string						`json:"province"`
	StreetName		string						`json:"streetName"`
	Map				string						`json:"map"`
	Images			[]string					`json:"images"`
	
}