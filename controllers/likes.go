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

func LikePost(c *fiber.Ctx) error{
	postID := c.Params("postId")
	userID := c.Params("userId")
	postObjID,err := primitive.ObjectIDFromHex(postID)
	if err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid post Id"})
	}
	userObjID,err := primitive.ObjectIDFromHex(userID)
	if err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid user Id"})
	}
	 var like models.LikesModel

	 like.UserID = userObjID
	 like.CreatedAt = time.Now().Unix()

	 collection := database.MongoCollection("posts")
	 _,err = collection.UpdateOne(
		context.Background(),
		bson.M{"_id":postObjID},
		bson.M{"$addToSet":bson.M{"likes":like}},
	 )
	 if err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Failed to like post"})
	 }
	 return c.Status(fiber.StatusOK).JSON(fiber.Map{"message":"Post Like Successfully"})
}


func UnlikePost(c *fiber.Ctx) error{
	postID := c.Params("postId")
    userID := c.Params("userId")
    postObjID, err := primitive.ObjectIDFromHex(postID)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid post ID"})
    }
    userObjID, err := primitive.ObjectIDFromHex(userID)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
    }

	collection := database.MongoCollection("posts")
	_,err = collection.UpdateOne(
		context.Background(),
		bson.M{"_id":postObjID},
		bson.M{"$pull":bson.M{"likes":bson.M{"userId":userObjID}}},
	)
	if err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Failed to unlike the post"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message":"Post unlike successfully"})
}