package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type Admin struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username    string `bson:"username" json:"username"`
	Password	string `bson:"password" json:"-"`
	Phone	   	string `bson:"phone" json:"phone"`
	IsVerified 	bool   `bson:"is_verified" json:"is_verified"`
	CreatedAt   primitive.DateTime `bson:"created_at" json:"created_at"`
	UpdatedAt   primitive.DateTime `bson:"updated_at" json:"updated_at"`
}