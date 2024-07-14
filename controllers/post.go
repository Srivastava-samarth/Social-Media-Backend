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

func CreatePost(c *fiber.Ctx) error{
	var post models.PostModel;
	if err := c.BodyParser(post); err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Cannot parse the post json"})
	}
	collection := database.MongoCollection("posts");
	post.ID = primitive.NewObjectID();
	post.CreatedAt = time.Now().Unix();
	post.UpdatedAt = time.Now().Unix();

	_,err := collection.InsertOne(context.Background(),post);
	if err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Failed to create the post"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message":"Post created successfully"})
}

func GetPosts(c *fiber.Ctx) error{
	ID:=c.Params("id");
	objID,err := primitive.ObjectIDFromHex(ID);
	if err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid post ID"})
	}
	var post models.PostModel
	collection := database.MongoCollection("posts")
	err = collection.FindOne(context.Background(),bson.M{"_id":objID}).Decode(&post);
	if err!=nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error":"Post not found"})
	}
	return c.JSON(post);
}

func UpdatePost(c *fiber.Ctx) error{
	ID := c.Params("id")
	objID,err := primitive.ObjectIDFromHex(ID)
	if err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid post id"})
	}
	var UpdatedData models.PostModel
	if err := c.BodyParser(UpdatedData);err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Error in parsing the data"})
	}

	UpdatedData.UpdatedAt = time.Now().Unix();
	collection := database.MongoCollection("posts");
	_,err = collection.UpdateOne(context.Background(),bson.M{"_id":objID},bson.M{"$set":UpdatedData},)
	if err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Cannot update the data"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message":"Post updated successfully"})
}

func DeletePost(c *fiber.Ctx) error{
	ID := c.Params("id");
	objID,err := primitive.ObjectIDFromHex(ID)
	if err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid post id"})
	}
	collection := database.MongoCollection("posts");
	_,err = collection.DeleteOne(context.Background(),bson.M{"_id":objID})
	if err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Failed to delete the post"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message":"Post deleted successfully"})
}