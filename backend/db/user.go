package db

import (
	"coffeh/model"
	"context"
	"log"
	"time"

	"github.com/go-telegram/bot/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (s *Store) CreateUser(ctx context.Context, user models.User) error {
	collection := s.client.Collection("users")
	userDocument := model.User{
		TelegramId: user.ID,
		Username:   user.Username,
		CreatedAt:  bson.NewDateTimeFromTime(time.Now()),
	}
	_, err := collection.InsertOne(ctx, userDocument)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return err
}

func (s *Store) FindUserByTelegramID(ctx context.Context, telegramId int64) (model.User, error) {
	collection := s.client.Collection("users")
	var res model.User
	filter := bson.D{{Key: "telegram_id", Value: telegramId}}

	err := collection.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println(err.Error())
			return model.User{}, err
		}
		log.Fatalf(err.Error())
	}

	return res, nil
}
