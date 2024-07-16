package controllers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/samarth-srivastava/social-media/database"
	"github.com/samarth-srivastava/social-media/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FollowUser(c *fiber.Ctx) error{
	followerID, err := primitive.ObjectIDFromHex(c.Params("followerID"))
	if err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid Follower Id"})
	}
	followingID,err := primitive.ObjectIDFromHex(c.Params("followingID"))
	if err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid following Id"})
	}
	collection := database.MongoCollection("follows");
	var follow models.FollowModel

	follow.ID =  primitive.NewObjectID()
	follow.FollowerID = followerID
	follow.FollowingID = followingID
	follow.CreatedAt = time.Now()

	_,err = collection.InsertOne(context.Background(),follow)
	if err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Cannot insert the follow data"})
	}

	userCollection := database.MongoCollection("users")
	_,err = userCollection.UpdateOne(
		context.Background(),
		bson.M{"_id":followerID},
		bson.M{"$inc":bson.M{"following":1}},
	)
	if err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Failed to update followers count"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message":"Successfully followed user"})
}

func UnfollowUser(c *fiber.Ctx) error{
	followerID,err := primitive.ObjectIDFromHex(c.Params("followerID"))
	if err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid Follower Id"})
	}
	followingID,err := primitive.ObjectIDFromHex(c.Params("followingID"))
	if err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid following Id"})
	}
	collection := database.MongoCollection("follows");
	_,err = collection.DeleteOne(context.Background(),bson.M{
		"followerId":followerID,
		"followingId":followingID,
	})
	if err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Cannot unfolllow the user"})
	}

	userCollection := database.MongoCollection("users")
	_,err = userCollection.UpdateOne(
		context.Background(),
		bson.M{"_id":followingID},
		bson.M{"$inc":bson.M{"followers":-1}},
	)
	if err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Failed to update the followers count"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message":"Successfully unfollow the user"})
}