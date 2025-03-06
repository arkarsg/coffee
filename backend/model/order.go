package model

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type OrderItem struct {
	ID    bson.ObjectID `bson:"_id"`
	Name  string        `bson:"item_name"`
	Price float32       `bson:"item_price"`
}

type Order struct {
	ID        bson.ObjectID `bson:"_id"`
	UserID    bson.ObjectID `bson:"user_id"`
	Quantity  int           `bson:"quantity"`
	Items     []OrderItem   `bson:"items"`
	createdAt bson.DateTime `bson:"createdAt"`
}
