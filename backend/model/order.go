package model

import (
	"coffeh/config"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type OrderItem struct {
	DrinkID    bson.ObjectID `bson:"drink_id" json:"drink_id" binding:"required"`
	Quantity   int           `bson:"quantity" json:"quantity" binding:"required,gt=0"`
	Variant    DrinkVariant  `bson:"drinkVariant" json:"drinkVariant" binding:"required,valid_drinkvariant"`
	IsTakeaway *bool         `bson:"isTakeaway" json:"isTakeaway" binding:"required"`
}

func (o *OrderItem) MaybeAddIced() float32 {
	if o.Variant == ICED {
		return config.ICED_SURCHARGE * float32(o.Quantity)
	}
	return float32(0)
}

func (o *OrderItem) MaybeAddTakeaway() float32 {
	if *o.IsTakeaway {
		return config.TAKEAWAY_SURCHARGE * float32(o.Quantity)
	}
	return float32(0)
}

type OrderStatus int16

const (
	PENDING_PAYMENT OrderStatus = iota
	PAID
	CANCELLED
)

type Order struct {
	ID             bson.ObjectID `bson:"_id" json:"id"`
	CustomerUserID int64         `bson:"customer_user_id" json:"customer_user_id"`
	Items          []OrderItem   `bson:"items" json:"items"`
	Total          float32       `bson:"total" json:"total"`
	IsPreorder     bool          `bson:"is_preorder" json:"is_preorder"`
	CollectFrom    bson.DateTime `bson:"collect_from" json:"collect_from"`
	CollectTo      bson.DateTime `bson:"collect_to" json:"collect_to"`
	OrderStatus    OrderStatus   `bson:"order_status" json:"order_status"`
	CreatedAt      bson.DateTime `bson:"createdAt" json:"createdAt"`
}

type CreateOrderDTO struct {
	CustomerUserID int64       `json:"customer_user_id" binding:"required"`
	Items          []OrderItem `json:"items" binding:"required,dive,required"`
	IsPreorder     *bool       `json:"is_preorder" binding:"required"`
	CollectFrom    time.Time   `json:"collect_from" binding:"required,ltefield=CollectTo" time_format:"2006-01-02 12:12:12"`
	CollectTo      time.Time   `json:"collect_to" binding:"required" time_format:"2006-01-02 12:12:12"`
}
