package db

import (
	"coffeh/model"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (s *Store) FindAllOrdersByTelegramId(telegramId int64) (model.Order, error) {
	collection := s.client.Collection("orders")
	var res model.Order
	filter := bson.D{{Key: "user_id", Value: telegramId}}

	err := collection.FindOne(context.TODO(), filter).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println(err.Error())
			return model.Order{}, err
		}
		log.Fatalf(err.Error())
	}

	return res, nil
}
