package admincontrollers

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
func AddAdmin(cReq *fiber.Ctx) error {
	var reqBody validators.CreateAdminValidator

	if err := cReq.BodyParser(&reqBody); err != nil {
		return utils.ResponseError(cReq, fiber.StatusBadRequest, "Bad request", nil)
	}
	if err := validators.Validate.Struct(reqBody); err != nil {
		return utils.ResponseError(cReq, fiber.StatusBadRequest, "Invalid request data", nil)
	}

	adminCollection := configs.GetCollection("admins")
	ctx, cancel := context.WithTimeout(cReq.Context(), 10*time.Second)
	defer cancel()

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(reqBody.Password), bcrypt.DefaultCost)
	if err != nil {
		return utils.ResponseError(cReq, fiber.StatusInternalServerError, "Failed to hash password", nil)
	}
	newAdmin := models.Admin{
		Username:    reqBody.Username,
		Password:    string(hashPassword),
		Phone:       reqBody.Phone,
		IsVerified:  true,
		CreatedAt:   primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt:   primitive.NewDateTimeFromTime(time.Now()),
	}
	result, err := adminCollection.InsertOne(ctx, newAdmin)
	if err != nil {
		cancel()
		return utils.ResponseError(cReq, fiber.StatusInternalServerError, "Failed to add admin", nil)
	}
	newAdmin.ID = result.InsertedID.(primitive.ObjectID)
	cancel()
	return utils.ResponseSuccess(cReq, fiber.StatusCreated, "Admin added successfully", newAdmin, nil)
}