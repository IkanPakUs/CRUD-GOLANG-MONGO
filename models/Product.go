package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID    primitive.ObjectID `bson:"_id" json:"id"`
	Name  string
	Stock int
	Price float64
}

type InputProduct struct {
	Name  string
	Stock int
	Price float64
}
