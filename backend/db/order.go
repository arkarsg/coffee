package db

import (
	"coffeh/model"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (s *Store) GetAllOrders(ctx context.Context) ([]model.Order, error) {
	collection := s.client.Collection("orders")
	cursor, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	res := []model.Order{}
	if err = cursor.All(ctx, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Store) FindAllOrdersByTelegramId(ctx context.Context, telegramId int64) ([]model.Order, error) {
	collection := s.client.Collection("orders")
	filter := bson.D{{Key: "customer_user_id", Value: telegramId}}

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

func (s *Store) CreateOrder(ctx context.Context, newOrder model.CreateOrderDTO) (bson.ObjectID, error) {
	collection := s.client.Collection("orders")

	total, err := s.CalculateTotalCost(ctx, newOrder.Items)
	if err != nil {
		log.Fatalf(err.Error())
		return bson.NilObjectID, err
	}

	order := model.Order{
		ID:             bson.NewObjectID(),
		CustomerUserID: newOrder.CustomerUserID,
		Items:          newOrder.Items,
		Total:          total,
		IsPreorder:     *newOrder.IsPreorder,
		CollectFrom:    bson.NewDateTimeFromTime(newOrder.CollectFrom),
		CollectTo:      bson.NewDateTimeFromTime(newOrder.CollectTo),
		OrderStatus:    model.PENDING_PAYMENT,
		CreatedAt:      bson.NewDateTimeFromTime(time.Now()),
	}

	id, err := collection.InsertOne(ctx, order)
	if err != nil {
		return bson.NilObjectID, err
	}
	return id.InsertedID.(bson.ObjectID), nil
}

func (s *Store) FulfillOrder(ctx context.Context, orderId bson.ObjectID) error {
	collection := s.client.Collection("orders")
	res := model.Order{}

	err := collection.FindOne(ctx, bson.D{{Key: "_id", Value: orderId}}).Decode(&res)
	if err != nil {
		return err
	}

	if res.OrderStatus == model.PAID {
		return fmt.Errorf("Order is already fulfilled")
	}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "order_status", Value: model.PAID}}}}
	_, err = collection.UpdateByID(ctx, orderId, update)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) CancelOrder(ctx context.Context, orderId bson.ObjectID) error {
	collection := s.client.Collection("orders")
	res := model.Order{}

	err := collection.FindOne(ctx, bson.D{{Key: "_id", Value: orderId}}).Decode(&res)
	if err != nil {
		return err
	}

	if res.OrderStatus == model.PAID {
		return fmt.Errorf("Order is already fulfilled")
	}

	if res.OrderStatus == model.CANCELLED {
		return fmt.Errorf("Order is already cancelled")
	}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "order_status", Value: model.CANCELLED}}}}
	_, err = collection.UpdateByID(ctx, orderId, update)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) DeleteOrder(ctx context.Context, orderId bson.ObjectID) error {
	collection := s.client.Collection("orders")
	res := model.Order{}

	err := collection.FindOne(ctx, bson.D{{Key: "_id", Value: orderId}}).Decode(&res)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(ctx, bson.D{{Key: "_id", Value: orderId}})
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) CalculateTotalCost(ctx context.Context, orderItems []model.OrderItem) (float32, error) {
	drinkCollection := s.client.Collection("drinks")
	totalCost := float32(0)

	for _, item := range orderItems {
		var drink model.Drink
		err := drinkCollection.FindOne(ctx, bson.M{"_id": item.DrinkID}).Decode(&drink)
		if err != nil {
			return 0, fmt.Errorf("failed to find drink price: %v", err)
		}

		totalCost += drink.Price*float32(item.Quantity) +
			item.MaybeAddTakeaway() +
			item.MaybeAddIced()
	}

	return totalCost, nil
}
