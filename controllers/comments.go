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

func AddComment(c *fiber.Ctx) error{
	postID := c.Params("postId")
	objID,err := primitive.ObjectIDFromHex(postID)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid post id"})
	}
	var comment models.CommentModel
	if err := c.BodyParser(comment); err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Cannot parse the comment"})
	}

	comment.ID = primitive.NewObjectID();
	comment.CreatedAt = time.Now().Unix();

	collection := database.MongoCollection("posts")
	_,err = collection.UpdateOne(context.Background(),bson.M{"_id":objID},bson.M{"$push":bson.M{"comments": comment}},)
	if err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Failed to add comment"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message":"Comment added successfully"})
}

func UpdateComment(c *fiber.Ctx) error{
	postID := c.Params("postId");
	commentID := c.Params("commentId")
	postObjID,err := primitive.ObjectIDFromHex(postID);
	if err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid Post Id"})
	}
	commentObjID,err := primitive.ObjectIDFromHex(commentID)
	if err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid comment Id"})
	}

	var updateData models.CommentModel;
	if err := c.BodyParser(updateData); err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Cannot parse the data"})
	}
	updateData.UpdatedAt = time.Now().Unix();

	collection := database.MongoCollection("posts")
	_, err = collection.UpdateOne(
        context.Background(),
        bson.M{"_id": postObjID, "comments._id": commentObjID},
        bson.M{"$set": bson.M{"comments.$.content": updateData.Content, "comments.$.updatedAt": updateData.UpdatedAt}},
    )
	if err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Failed to update comment"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message":"Comment updated successfully"})
}

func DeleteComment(c *fiber.Ctx) error{
	postID := c.Params("postId");
	commentID := c.Params("commentId")
	postObjID,err := primitive.ObjectIDFromHex(postID);
	if err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid Post Id"})
	}
	commentObjID,err := primitive.ObjectIDFromHex(commentID)
	if err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid comment Id"})
	}
	collection := database.MongoCollection("posts")
	_,err = collection.UpdateOne(
		context.Background(),
		bson.M{"_id":postObjID},
		bson.M{"$pull":bson.M{"comments":bson.M{"_id":commentObjID}}},
	)
	if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete comment"})
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Comment deleted successfully"})
}