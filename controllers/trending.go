package controllers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/samarth-srivastava/social-media/database"
	"github.com/samarth-srivastava/social-media/models"
	"go.mongodb.org/mongo-driver/bson"
)

type TrendingPost struct {
	models.PostModel
	EngagemantScore int `json:"engagementScore"`
}

func GetTrendingPosts(c *fiber.Ctx) error{
	collection := database.MongoCollection("posts")

	timeWindow := time.Now().AddDate(0,0,-7);

	filter := []bson.M{
		{
			"$match":bson.M{
				"$or":[]bson.M{
					{"createdAt":bson.M{"$gte":timeWindow}},
					{"likes.createdAt":bson.M{"$gte": timeWindow}},
					{"comments.createdAt":bson.M{"$gte": timeWindow}},
				},
			},
		},
		{
			"$project": bson.M{
                "content":       1,
                "userId":        1,
                "createdAt":     1,
                "likes":         1,
                "comments":      1,
                "likesCount":    bson.M{"$size": "$likes"},
                "commentsCount": bson.M{"$size": "$comments"},
            },
		},
		{
            "$addFields": bson.M{
                "engagementScore": bson.M{
                    "$add": []interface{}{
                        bson.M{"$multiply": []interface{}{"$likesCount", 2}},
                        bson.M{"$multiply": []interface{}{"$commentsCount", 1}},
                    },
                },
            },
        },
		{
            "$sort": bson.M{"engagementScore": -1, "createdAt": -1},
        },
        {
            "$limit": 10, // Limit to top 10 trending posts
        },
	}
	ctx, cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel();

	cursor, err := collection.Aggregate(ctx,filter);
	if err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Failed to get trending posts"})
	}
	defer cursor.Close(ctx);

	var trendingPosts []TrendingPost
    if err := cursor.All(ctx, &trendingPosts); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to decode trending posts"})
    }

    return c.Status(fiber.StatusOK).JSON(trendingPosts)
}