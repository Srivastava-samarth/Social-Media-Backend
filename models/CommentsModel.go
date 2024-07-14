package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type CommentModel struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID    primitive.ObjectID `json:"userId,omitempty" bson:"userId,omitempty"`
	Content   string             `json:"content,omitempty" bson:"content,omitempty"`
	CreatedAt int64              `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt int64              `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
