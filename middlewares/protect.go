package middlewares

import (
	"ELEVATE_INVIX_BE/configs"
	"ELEVATE_INVIX_BE/utils"
	"context"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UserProtect(cReq *fiber.Ctx) error {
	authHeader := cReq.Get("Authorization")
	if authHeader == ""  || !strings.HasPrefix(authHeader, "Bearer ") {
		return utils.ResponseError(cReq, fiber.StatusUnauthorized, "Unauthorized", nil)
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := utils.VerifyToken(token)
	if err != nil {
		return utils.ResponseError(cReq, fiber.StatusUnauthorized, "Unauthorized", nil)
	}

	userIDstr, ok := claims["user_id"].(string)
	if !ok {
		return utils.ResponseError(cReq, fiber.StatusUnauthorized, "Invalid token payload", nil)
	}
	userID, err := primitive.ObjectIDFromHex(userIDstr)
	if err != nil {
		return utils.ResponseError(cReq, fiber.StatusUnauthorized, "Invalid user ID", nil)
	}

	sessionCollection := configs.GetCollection("sessions")
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	filter := bson.M{
		"user_id": userID,
		"is_active": true,
		"token": token,
	}
	count, err := sessionCollection.CountDocuments(ctx, filter)
	if err != nil {
		return utils.ResponseError(cReq, fiber.StatusInternalServerError, "Failed to check session", nil)
	}
	if count == 0 {
		return utils.ResponseError(cReq, fiber.StatusUnauthorized, "Session not found or inactive", nil)
	}
	cReq.Locals("userID", userID)
	cReq.Locals("token", token)

	return cReq.Next()
}
