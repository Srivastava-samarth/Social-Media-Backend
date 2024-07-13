
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PostModel struct {
    ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Content      string             `json:"content,omitempty" bson:"content,omitempty"`
    CreatedAt    int64              `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
    UpdatedAt    int64              `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
    AuthorID     primitive.ObjectID `json:"authorId,omitempty" bson:"authorId,omitempty"`
    LikeCount    int                `json:"likeCount,omitempty" bson:"likeCount,omitempty"`
    CommentCount int                `json:"commentCount,omitempty" bson:"commentCount,omitempty"`
}
