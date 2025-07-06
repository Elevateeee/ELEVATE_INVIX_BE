package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Sesion struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	CreatedAt primitive.DateTime `bson:"created_at" json:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at" json:"updated_at"`
	ExpiresAt primitive.DateTime `bson:"expires_at" json:"expires_at"`
	Token 	  string 			 `bson:"token,omitempty" json:"token"`
	IsActive  bool               `bson:"is_active" json:"is_active"`
}
