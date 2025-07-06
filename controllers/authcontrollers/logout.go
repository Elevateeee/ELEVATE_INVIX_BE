package authcontrollers

import (
	"ELEVATE_INVIX_BE/configs"
	"ELEVATE_INVIX_BE/utils"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)



func Logout(cReq *fiber.Ctx) error {
	userID := cReq.Locals("userID").(primitive.ObjectID)
	token := cReq.Locals("token").(string)

	sessionCollection := configs.GetCollection("sessions")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"user_id": userID,
		"token":   token,
		"is_active": true,
	}
	result, err := sessionCollection.DeleteOne(ctx, filter)
	if err != nil {
		return utils.ResponseError(cReq, fiber.StatusInternalServerError, "Failed to logout", nil)
	}
	if result.DeletedCount == 0 {
		return utils.ResponseError(cReq, fiber.StatusNotFound, "Session not found", nil)
	}

	return utils.ResponseSuccess(cReq, fiber.StatusOK, "Logout successful", nil, nil)
}
