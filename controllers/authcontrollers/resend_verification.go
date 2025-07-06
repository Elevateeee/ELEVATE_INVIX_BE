package authcontrollers

import (
	"ELEVATE_INVIX_BE/configs"
	"ELEVATE_INVIX_BE/models"
	"ELEVATE_INVIX_BE/utils"
	"ELEVATE_INVIX_BE/validators"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)


func ResendVerification(cReq *fiber.Ctx) error {
	var reqbody validators.ResendVerificationValidator

	if err := cReq.BodyParser(&reqbody); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Bad request")
	}
	if err := validators.Validate.Struct(reqbody); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request data")
	}

	userCollection := configs.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var userData models.User
	err := userCollection.FindOne(ctx, bson.M{"email": reqbody.Email}).Decode(&userData)
	if err != nil {
		return utils.ResponseError(cReq, fiber.StatusNotFound, "User not found", nil)
	}

	if userData.IsVerified {
		return utils.ResponseError(cReq, fiber.StatusBadRequest, "Email is already verified", nil)
	}

	newToken, err := utils.GenerateToken(userData.ID.Hex(), userData.Email)
	if err != nil {
		log.Println("Failed to generate verification token:", err)
		return utils.ResponseError(cReq, fiber.StatusInternalServerError, "Failed to generate verification token", nil)
	}

	frontendURL := os.Getenv("FE_URL")
	verificationLink := fmt.Sprintf("%s/verify-email?token=%s", frontendURL, newToken)

	err = utils.SendVerificationEmail(userData.Email, userData.Username, verificationLink)
	if err != nil {
		log.Println("Failed to send verification email:", err)
		return utils.ResponseError(cReq, fiber.StatusInternalServerError, "Failed to send verification email", nil)
	}

	return utils.ResponseSuccess(cReq, fiber.StatusOK, "Verification email resent successfully", nil, nil)
}
