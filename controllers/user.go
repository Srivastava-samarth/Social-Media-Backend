package controllers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/samarth-srivastava/social-media/database"
	"github.com/samarth-srivastava/social-media/middleware"
	"github.com/samarth-srivastava/social-media/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string,error){
	bytes,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	return string(bytes),err;
}

func Register(c *fiber.Ctx) error{
	var user models.UserModel;
	if err := c.BodyParser(user); err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Cannot parse the JSON"})
	}
	collection := database.MongoCollection("users");
	filter := bson.M{
		"$or":[]bson.M{
			{"username":user.Username},
			{"email":user.Email},
		},
	}
	var existingUser models.UserModel
	err := collection.FindOne(context.Background(),filter).Decode(&existingUser)
	if err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User with given username or email already exists"})
	}
	hashedPassword, err := hashPassword(user.Password)
	if err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Failed to hash the password"})
	}
	user.Password = hashedPassword;
	user.ID = primitive.NewObjectID();
	user.CreatedAt = time.Now().Unix();
	user.UpdatedAt = time.Now().Unix();

	_,err = collection.InsertOne(context.Background(),user);
	if err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Failed to add the user to the database"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message":"User registered successfully"})
}

func Login(c *fiber.Ctx) error{
	var user models.LoginUser;
	if err := c.BodyParser(user);
	 err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Cannot parse the user data"})
	}

	collection := database.MongoCollection("users");
	filter := bson.M{"email":user.Email}
	var existingUser models.UserModel;
	err := collection.FindOne(context.Background(),filter).Decode(&existingUser)
	if err!=nil{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error":"Invalid email or password"})
	}
	 err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password),[]byte(user.Password))
	 if err!=nil{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error":"Invalid email or password"})
	 }

	 token, err := middleware.GenerateJWT(existingUser.ID.Hex())
	 if err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Failed to generate token"})
	 }
	 return c.JSON(fiber.Map{"token":token});
}