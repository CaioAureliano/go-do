package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Task      string             `bson:"task" json:"task"`
	Status    bool               `bson:"status" json:"status,omitempty"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at,omitempty"`
}
