package usercontrollers

import (
	"ELEVATE_INVIX_BE/configs"
	"ELEVATE_INVIX_BE/utils"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ListUsers(cReq *fiber.Ctx) error {
	params := utils.ParseQueryParams(cReq)

	request := bson.M{}
	if params.Search != "" {
		request["$or"] = []bson.M{
			{"username": bson.M{"$regex": params.Search, "$options": "i"}},
			{"email": bson.M{"$regex": params.Search, "$options": "i"}},
			{"phone": bson.M{"$regex": params.Search, "$options": "i"}},
		}
	}

	skip := (params.Page - 1) * params.Limit
	options := options.Find().
		SetSort(bson.D{{Key: params.Sort, Value: params.Order}}).
		SetSkip(int64(skip)).
		SetLimit(int64(params.Limit))
		
	userCollection := configs.GetCollection("users")
	ctx, cancel := context.WithTimeout(cReq.Context(), 10 *time.Second)
	defer cancel()

	usersData, err := userCollection.Find(ctx, request, options)
	if err != nil {
		return utils.ResponseError(cReq, fiber.StatusInternalServerError, "Failed to fetch users", nil)
	}
	defer usersData.Close(ctx) 

	var listUsers []bson.M
	if err := usersData.All(ctx, &listUsers); err != nil {
		return utils.ResponseError(cReq, fiber.StatusInternalServerError, "Failed to parse users data", nil)
	}
	totalData, err := userCollection.CountDocuments(ctx, request)
	if err != nil {
		return utils.ResponseError(cReq, fiber.StatusInternalServerError, "Failed to count users", nil)
	}
	pagination := &utils.Pagination{
		Total : int(totalData),
		Page:  params.Page,
		Limit: params.Limit,
		Pages: int((totalData + int64(params.Limit) - 1) / int64(params.Limit)),
	}
	return utils.ResponseSuccess(cReq, fiber.StatusOK, "List of users", listUsers, pagination)
}
