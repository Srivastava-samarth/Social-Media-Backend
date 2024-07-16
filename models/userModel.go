package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserModel struct {
	ID        primitive.ObjectID `json: "id,omitempty" bson:"_id, omitempty"`
	Username  string             `json:"username,omitempty" bson: "username,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	Password  string             `json:"password,omitempty" bson:"password,omitempty"`
	Profile   UserProfile        `json:"profile,omitempty" bson:"profile,omitempty"`
	CreatedAt int64              `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt int64              `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	Followers int                `json:"followers" bson:"followers"`
	Following int                `json:"following" bson:"following"`
}

type UserProfile struct {
	Bio       string `json:"bio,omitempty" bson:"bio,omitempty"`
	AvatarUrl string `json:"avatarUrl,omitempty" bson:"avatarUrl,omitempty"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
