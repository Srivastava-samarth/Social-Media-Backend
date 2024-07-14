package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PostModel struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID    primitive.ObjectID `json:"userId,omitempty" bson:"userId,omitempty"`
	Content   string             `json:"content,omitempty" bson:"content,omitempty"`
	CreatedAt int64              `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt int64              `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	AuthorID  primitive.ObjectID `json:"authorId,omitempty" bson:"authorId,omitempty"`
	Likes     []LikesModel       `json:"likes,omitempty" bson:"likes,omitempty"`
	Comment   []CommentModel     `json:"comment,omitempty" bson:"comment,omitempty"`
}
