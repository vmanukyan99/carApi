package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Car struct {
	ID      primitive.ObjectID `json:"_id" bson:"_id"`
	Brand   string             `json:"brand" bson:"brand"`
	Model   string             `json:"model" bson:"model"`
	Price   uint               `json:"price" bson:"price"`
	Status  string             `json:"status" bson:"status"`
	Mileage uint               `json:"mileage" bson:"mileage"`
}
