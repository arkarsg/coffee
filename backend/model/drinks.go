package model

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type DrinkCategory string

const (
	COFFEE     DrinkCategory = "caffeinated"
	NON_COFFEE DrinkCategory = "non-caffeinated"
)

type DrinkVariant string

const (
	ICED DrinkVariant = "iced"
	HOT  DrinkVariant = "hot"
)

type Drink struct {
	ID          bson.ObjectID  `bson:"_id" json:"id"`
	Name        string         `bson:"name" json:"name"`
	Slug        string         `bson:"slug" json:"slug"`
	Price       float32        `bson:"price" json:"price"`
	Category    DrinkCategory  `bson:"category" json:"category"`
	Tags        []string       `bson:"tags" json:"tags"`
	Variants    []DrinkVariant `bson:"drinkVariants" json:"drinkVariants"`
	Description string         `bson:"description" json:"description"`
	CreatedAt   bson.DateTime  `bson:"createdAt" json:"createdAt"`
}

type CreateDrinkDTO struct {
	Name        string         `bson:"name" json:"name" binding:"required,max=30"`
	Price       float32        `bson:"price" json:"price" binding:"required,gt=0"`
	Category    DrinkCategory  `bson:"category" json:"category" binding:"required,valid_drinkcategory"`
	Tags        []string       `bson:"tags" json:"tags"`
	Variants    []DrinkVariant `bson:"drinkVariants" json:"drinkVariants" binding:"required,dive,required,valid_drinkvariant"`
	Description string         `bson:"description" json:"description" binding:"required,max=100"`
}

type UpdateDrinkDTO struct {
	Price       float32        `bson:"price,omitempty" json:"price,omitempty" binding:"omitempty,gt=0"`
	Category    DrinkCategory  `bson:"category,omitempty" json:"category,omitempty" binding:"omitempty,valid_drinkcategory"`
	Tags        []string       `bson:"tags" json:"tags"`
	Variants    []DrinkVariant `bson:"drinkVariants,omitempty" json:"drinkVariants,omitempty" binding:"dive,valid_drinkvariant,omitempty"`
	Description string         `bson:"description,omitempty" json:"description,omitempty" binding:"max=100"`
}
