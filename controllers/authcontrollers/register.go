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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(cReq *fiber.Ctx) error { 

	var input validators.RegisterValidator

	if err := cReq.BodyParser(&input); err != nil {
		return utils.ResponseError(cReq, fiber.StatusBadRequest, "Bad request", nil)
	}
	if err := validators.Validate.Struct(input); err != nil {
		return utils.ResponseError(cReq, fiber.StatusBadRequest, "Invalid request data",nil )
	}

	userCollection := configs.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10 *time.Second)
	defer cancel()

	if existEmail, _ := utils.IsFieldExists(ctx, userCollection, "email", input.Email); existEmail {
		return utils.ResponseError(cReq, fiber.StatusConflict, "Email already exists", nil)
	}
	if existUsername, _ := utils.IsFieldExists(ctx, userCollection, "username", input.Username); existUsername {
		return utils.ResponseError(cReq, fiber.StatusConflict, "Username already exists", nil)
	}
	if existPhone, _ := utils.IsFieldExists(ctx, userCollection, "phone", input.Phone); existPhone {
		return utils.ResponseError(cReq, fiber.StatusConflict, "Phone number already exists", nil)
	}

	
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return utils.ResponseError(cReq, fiber.StatusInternalServerError, "Failed to hash password", nil)
	}

	newUser := models.User{
		Username: input.Username,
		Email:  input.Email,
		Password: string(hashPassword),
		IsVerified: false,
		Phone: input.Phone,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}

	
	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		log.Println("Error inserting new user:", err)
		return utils.ResponseError(cReq, fiber.StatusInternalServerError, "Failed to create user", nil)
	}

	newUser.ID = result.InsertedID.(primitive.ObjectID)

	token, err := utils.GenerateToken(newUser.ID.Hex(), newUser.Email)
	if err != nil {
		log.Println("Gagal generate token:", err)
		return utils.ResponseError(cReq, fiber.StatusInternalServerError, "Failed to generate verification token", nil)
	}

	frontendURL := os.Getenv("FE_URL")
	verificationLink := fmt.Sprintf("%s/verify-email?token=%s", frontendURL, token)

	fmt.Println("Verification Link:", verificationLink)

	err = utils.SendVerificationEmail(newUser.Email, newUser.Username, verificationLink)
	if err != nil {
		log.Println("Gagal kirim email verifikasi:", err)
	}

	return utils.ResponseSuccess(cReq, fiber.StatusCreated, "User registered successfully", fiber.Map{
		"username":   newUser.Username,
		"created_at": newUser.CreatedAt,
	}, nil)
}
