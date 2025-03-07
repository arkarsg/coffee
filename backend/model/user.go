package model

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID         bson.ObjectID `bson:"_id" json:"id"`
	TelegramId int64         `bson:"telegram_id" json:"telegram_id"`
	Username   string        `bson:"username" json:"username"`
	CreatedAt  bson.DateTime `bson:"createdAt" json:"createdAt"`
}
