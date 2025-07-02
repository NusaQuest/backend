package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Proposal struct {
	Id                  primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ScId                int64              `json:"scId"`
	ScTargets           []string           `json:"scTargets"`
	ScValues            []int64            `json:"scValues"`
	ScCalldatas         []string           `json:"scCalldatas"`
	Wallet              string             `json:"wallet"`
	ProposalName        string             `json:"name" validate:"required"`
	ProposalDescription string             `json:"proposalDescription" validate:"required,min=8,max=256"`
	BeachName           string             `json:"beachName" validate:"required"`
	City                string             `json:"city" validate:"required"`
	Province            string             `json:"province" validate:"required"`
	Map                 string             `json:"map" validate:"required"`
	Images              []string           `json:"images" validate:"required"`
}
