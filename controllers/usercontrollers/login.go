package usercontrollers

import (
	"ELEVATE_INVIX_BE/configs"
	"ELEVATE_INVIX_BE/models"
	"ELEVATE_INVIX_BE/utils"
	"ELEVATE_INVIX_BE/validators"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
)

func Login(cReq *fiber.Ctx) error {
	var reqBody validators.LoginValidator

	if err := cReq.BodyParser(&reqBody); err != nil {
		return utils.ResponseError(cReq, fiber.StatusBadRequest, "Bad request", nil)
	}
	if err := validators.Validate.Struct(reqBody); err != nil {
		return utils.ResponseError(cReq, fiber.StatusBadRequest, "Invalid request data", nil)
	}

	userCollection := configs.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var userData bson.M
	err := userCollection.FindOne(ctx, bson.M{"email": reqBody.Email}).Decode(&userData)
	if err != nil {
		return utils.ResponseError(cReq, fiber.StatusUnauthorized, "Invalid email or password", nil)
	}

	userID := userData["_id"].(primitive.ObjectID)

	sessionCollection := configs.GetCollection("sessions")
	sessionFilter := bson.M{
		"user_id":  userID,
		"is_active": true,
	}
	count, err := sessionCollection.CountDocuments(ctx, sessionFilter)
	if err != nil {
		return utils.ResponseError(cReq, fiber.StatusInternalServerError, "Failed to check session", nil)
	}
	if count > 0 {
		return utils.ResponseError(cReq, fiber.StatusConflict, "User already logged in", nil)
	}

	token, err := utils.GenerateToken(userID.Hex(), userData["email"].(string))
	if err != nil {
		return utils.ResponseError(cReq, fiber.StatusInternalServerError, "Failed to generate token", nil)
	}

	newSession := models.Sesion{
		UserID:    userID,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
		ExpiresAt: primitive.NewDateTimeFromTime(time.Now().Add(24 * time.Hour)),
		Token:     token,
		IsActive:  true,
	}

	result, err := sessionCollection.InsertOne(ctx, newSession)
	if err != nil {
		return utils.ResponseError(cReq, fiber.StatusInternalServerError, "Failed to create session", nil)
	}
	newSession.ID = result.InsertedID.(primitive.ObjectID)

	return utils.ResponseSuccess(cReq, fiber.StatusOK, "Login successful", fiber.Map{
		"token": newSession.Token,
	}, nil)
}
