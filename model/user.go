package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name string `bson:"name" json:"name"`
	Email string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}

type UserCreateDto struct {
	Name     string `bson:"name" json:"name" binding:"required,min=2,max=50" validate:"required,min=2,max=50" customError:"Name is required and must be between 2 and 50 characters."`
	Email    string `bson:"email" json:"email" binding:"required,email" validate:"required,email" customError:"Please provide a valid email address."`
	Password string `bson:"password" json:"password" binding:"required,min=6" validate:"required,min=6" customError:"Password must be at least 6 characters long."`
}