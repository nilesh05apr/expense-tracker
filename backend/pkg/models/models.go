package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Expense struct {
	ID 					primitive.ObjectID 	`json:"_id,omitempty" bson:"_id,omitempty"`
	Amount 				float64 		  	`json:"amount,omitempty" bson:"amount,omitempty"`
	Description 		string 	  			`json:"description,omitempty" bson:"description,omitempty"`
	Note 				string 		  		`json:"note,omitempty" bson:"note,omitempty"`
	CreatedAt 			time.Time   		`json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	ExpenseCategoryID 	primitive.ObjectID 	`json:"expenseCategoryId,omitempty" bson:"expenseCategoryId,omitempty"`
}

type ExpenseCategory struct {
	ID 			primitive.ObjectID 		`json:"_id,omitempty" bson:"_id,omitempty"`
	Name 		string 		  			`json:"name,omitempty" bson:"name,omitempty"`
	Amount 		float64 		  		`json:"amount,omitempty" bson:"amount,omitempty"`
	MaxAmount 	float64 	  			`json:"maxAmount,omitempty" bson:"maxAmount,omitempty"`
}