package middleware

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func Middleware(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing JWT"})
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil,fiber.ErrUnauthorized;
		}
		return []byte(os.Getenv("JWT_SECRET")),nil;
	})

	if err != nil || !token.Valid{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error":"Invalid Token"})
	}
	claims,ok := token.Claims.(jwt.MapClaims);
	if !ok || !token.Valid{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"eeror":"Invalid jwt claims"})
	}
	c.Locals("userID",claims["user_id"]);
	c.Locals("exp",claims["exp"]);
	return c.Next();
}

func GenerateJWT(userID string) (string,error){
	claims := jwt.MapClaims{
		"user_id":userID,
		"exp" : time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims);
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")));
}

func LoggerMiddleware(c *fiber.Ctx) error{
	start := time.Now();
	err := c.Next();
	duration := time.Since(start);

	log.Println("%s - %s %s %s - %s",time.Now().Format(time.RFC3339),c.Method(),c.Path(),c.IP(),duration)
	return err;
}

func ErrorHandlerMiddlerware(c *fiber.Ctx,err error) error{
	if err != nil{
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
		return nil;
	}
	return c.Next();
}