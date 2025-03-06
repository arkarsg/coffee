package db

import (
	"coffeh/model"
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (s *Store) GetAllOrders(ctx context.Context) ([]model.Order, error) {
	collection := s.client.Collection("orders")
	cursor, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	res := []model.Order{}
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Store) FindAllOrdersByTelegramId(ctx context.Context, telegramId int64) ([]model.Order, error) {
	collection := s.client.Collection("orders")
	filter := bson.D{{Key: "user_id", Value: telegramId}}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	res := []model.Order{}
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}
