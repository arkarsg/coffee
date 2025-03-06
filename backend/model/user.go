package model

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID         bson.ObjectID `bson:"_id"`
	TelegramId int64         `bson:"telegram_id"`
	Username   string        `bson:"username"`
	IsAdmin    bool          `bson:"is_admin"`
	CreatedAt  bson.DateTime `bson:"createdAt"`
}
