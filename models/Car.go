package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Car struct {
	ID              primitive.ObjectID `bson:"_id"`
	Name            *string            `json:"name" validate:"required,min=2,max=100"`
	Price           *float64           `json:"price" validate:"required"`
	User_id         *string            `json:"user_id" validate:"required"`
	Created_at      time.Time          `json:"created_at"`
	Updated_at      time.Time          `json:"updated_at"`
	State           string             `json:"state"`
	Status          string             `json:"status"`
	Manufacturer_id *string            `json:"manufacturer_id"`
	Model           string             `json:"model"`
	Body_type       string             `json:"body_type"`
}
