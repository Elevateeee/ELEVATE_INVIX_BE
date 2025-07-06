package utils

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func IsFieldExists(ctx context.Context, collection *mongo.Collection, field string, value interface{}) (bool, error) {
	filter := bson.M{field: value}
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
