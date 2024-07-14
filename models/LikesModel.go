package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type LikesModel struct {
	UserID    primitive.ObjectID `json:"userId,omitempty" bson:"userId,omitempty"`
	CreatedAt int64              `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}