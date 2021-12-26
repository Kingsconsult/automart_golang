package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Manufacturer struct {
	ID              primitive.ObjectID `bson:"_id"`
	Name            *string            `json:"name" validate:"required,min=2,max=100"`
	Created_at      time.Time          `json:"created_at"`
	Updated_at      time.Time          `json:"updated_at"`
	Manufacturer_id string             `json:"car_id"`
}
