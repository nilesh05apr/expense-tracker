package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExpenseCategory struct {
	ID 			primitive.ObjectID 		`json:"_id,omitempty" bson:"_id,omitempty"`
	Name 		string 		  			`json:"name,omitempty" bson:"name,omitempty"`
	Amount 		float64 		  		`json:"amount,omitempty" bson:"amount,omitempty" default:"0"`
	MaxAmount 	float64 	  			`json:"maxAmount,omitempty" bson:"maxAmount,omitempty"`
}