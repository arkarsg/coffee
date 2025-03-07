package db

import (
	"coffeh/model"
	"context"
	"log"
	"time"

	"github.com/gosimple/slug"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (s *Store) CreateDrink(ctx context.Context, newDrink model.CreateDrinkDTO) error {
	collection := s.client.Collection("drinks")
	userDocument := model.Drink{
		ID:          bson.NewObjectID(),
		Name:        newDrink.Name,
		Slug:        slug.Make(newDrink.Name),
		Price:       newDrink.Price,
		Category:    newDrink.Category,
		Tags:        newDrink.Tags,
		Variants:    newDrink.Variants,
		Description: newDrink.Description,
		CreatedAt:   bson.NewDateTimeFromTime(time.Now()),
	}
	_, err := collection.InsertOne(ctx, userDocument)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return err
}

func (s *Store) GetAllDrinks(ctx context.Context) ([]model.Drink, error) {
	collection := s.client.Collection("drinks")
	cursor, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	res := []model.Drink{}
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Store) GetDrink(ctx context.Context, drinkSlug string) (model.Drink, error) {
	collection := s.client.Collection("drinks")
	res := model.Drink{}

	err := collection.FindOne(ctx, bson.D{{Key: "slug", Value: drinkSlug}}).Decode(&res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Store) UpdateDrinkBySlug(ctx context.Context, drinkSlug string, updatedDrink model.UpdateDrinkDTO) error {
	collection := s.client.Collection("drinks")
	res := model.Drink{}

	err := collection.FindOne(ctx, bson.D{{Key: "slug", Value: drinkSlug}}).Decode(&res)
	if err != nil {
		return err
	}
	// update only if the values are provided
	updateFields := bson.D{}

	if updatedDrink.Price != 0 {
		updateFields = append(updateFields, bson.E{Key: "price", Value: updatedDrink.Price})
	}

	if updatedDrink.Category != "" {
		updateFields = append(updateFields, bson.E{Key: "category", Value: updatedDrink.Category})
	}

	if updatedDrink.Tags != nil {
		updateFields = append(updateFields, bson.E{Key: "tags", Value: updatedDrink.Tags})
	}

	if updatedDrink.Variants != nil {
		updateFields = append(updateFields, bson.E{Key: "drinkVariants", Value: updatedDrink.Variants})
	}
	if updatedDrink.Description != "" {
		updateFields = append(updateFields, bson.E{Key: "description", Value: updatedDrink.Description})
	}

	if len(updateFields) == 0 {
		return nil
	}

	filter := bson.D{{Key: "slug", Value: drinkSlug}}
	update := bson.D{{Key: "$set", Value: updateFields}}
	_, err = collection.UpdateOne(ctx, filter, update)
	return err
}

func (s *Store) DeleteDrink(ctx context.Context, drinkSlug string) error {
	collection := s.client.Collection("drinks")
	res := model.Drink{}

	err := collection.FindOne(ctx, bson.D{{Key: "slug", Value: drinkSlug}}).Decode(&res)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(ctx, bson.D{{Key: "slug", Value: drinkSlug}})
	if err != nil {
		return err
	}
	return nil
}
