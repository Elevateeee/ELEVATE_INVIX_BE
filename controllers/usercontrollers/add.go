package usercontrollers

import (
	"ELEVATE_INVIX_BE/configs"
	"ELEVATE_INVIX_BE/models"
	"ELEVATE_INVIX_BE/utils"
	"ELEVATE_INVIX_BE/validators"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func AddUser (cReq *fiber.Ctx) error {
	var reqBody validators.CreateUserValidator

	if err := cReq.BodyParser(&reqBody); err != nil {
		return utils.ResponseError(cReq, fiber.StatusBadRequest, "Bad request", nil)
	}
	if err := validators.Validate.Struct(reqBody); err != nil {
		return utils.ResponseError(cReq, fiber.StatusBadRequest, "Invalid request data", nil)
	}
	
	userCollection := configs.GetCollection("users")
	ctx, cancle := context.WithTimeout(cReq.Context(), 10 * time.Second)
	defer cancle()

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(reqBody.Password), bcrypt.DefaultCost)
	if err != nil {
		return utils.ResponseError(cReq, fiber.StatusInternalServerError, "Failed to hash password", nil)
	}
	newUser := models.User{
		Username:   reqBody.Username,
		Email:      reqBody.Email,
		Password:   string(hashPassword),
		IsVerified: true,
		Phone:      reqBody.Phone,
		CreatedAt:  primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt:  primitive.NewDateTimeFromTime(time.Now()),
	}
	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		cancle()
		return utils.ResponseError(cReq, fiber.StatusInternalServerError, "Failed to add user", nil)
	}
	newUser.ID = result.InsertedID.(primitive.ObjectID)
	cancle()
	return utils.ResponseSuccess(cReq, fiber.StatusCreated, "User added successfully", newUser, nil)
}
