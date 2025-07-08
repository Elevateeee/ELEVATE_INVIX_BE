package usercontrollers

import (
	"ELEVATE_INVIX_BE/configs"
	"ELEVATE_INVIX_BE/utils"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func VerifyEmail(cReq *fiber.Ctx) error {
	token := cReq.Query("token")
	if token == "" {
		return utils.ResponseError(cReq, fiber.StatusBadRequest, "Bad request", nil)
	}
	claims, err := utils.VerifyToken(token)
	if err != nil {
		return utils.ResponseError(cReq, fiber.StatusBadRequest, "Invalid or expired token", nil)
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return utils.ResponseError(cReq, fiber.StatusBadRequest, "Invalid token payload", nil)
	}
	email, ok := claims["email"].(string)
	if !ok {
		return utils.ResponseError(cReq, fiber.StatusBadRequest, "Invalid token payload", nil)
	}

	userCollection := configs.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return utils.ResponseError(cReq, fiber.StatusBadRequest, "Invalid user ID", err)
	}

	var existingUser bson.M
	err = userCollection.FindOne(ctx, bson.M{"_id": objectID, "email": email}).Decode(&existingUser)
	if err != nil {
		return utils.ResponseError(cReq, fiber.StatusNotFound, "User not found", nil)
	}

	if isVerified, ok := existingUser["is_verified"].(bool); ok && isVerified {
		return utils.ResponseError(cReq, fiber.StatusBadRequest, "Email is already verified", nil)
	}
	
	update := bson.M{"$set": bson.M{
		"is_verified": true,
		"updated_at":  primitive.NewDateTimeFromTime(time.Now()),
	}}

	_, err = userCollection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return utils.ResponseError(cReq, fiber.StatusInternalServerError, "Failed to update user data", nil)
	}

	return utils.ResponseSuccess(cReq, fiber.StatusOK, "Email verified successfully", nil, nil)
}
