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

func (s *Store) CreateUser(user models.User) error {
	collection := s.client.Collection("users")
	userDocument := model.User{
		ID:         bson.NewObjectID(),
		TelegramId: user.ID,
		Username:   user.Username,
		IsAdmin:    false,
		CreatedAt:  bson.NewDateTimeFromTime(time.Now()),
	}
	_, err := collection.InsertOne(context.TODO(), userDocument)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return err
}

func (s *Store) FindUserByTelegramID(telegramId int64) (model.User, error) {
	collection := s.client.Collection("users")
	var res model.User
	filter := bson.D{{Key: "telegram_id", Value: telegramId}}

	err := collection.FindOne(context.TODO(), filter).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println(err.Error())
			return model.User{}, err
		}
		log.Fatalf(err.Error())
	}

	return res, nil
}

func (s *Store) IsAdmin(telegramId int64) (bool, error) {
	collection := s.client.Collection("users")
	var res model.User
	filter := bson.D{{Key: "telegram_id", Value: telegramId}}

	err := collection.FindOne(context.TODO(), filter).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println(err.Error())
			return false, err
		}
		log.Fatalf(err.Error())
	}

	return res.IsAdmin, nil
}
