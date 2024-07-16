package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FollowModel struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FollowerID  primitive.ObjectID `json:"followerId" bson:"followerId"`
	FollowingID primitive.ObjectID `json:"followingId" bson:"followingId"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
}
