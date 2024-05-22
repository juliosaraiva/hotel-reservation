package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Email             string             `bson:"email" json:"email"`
	FirstName         string             `bson:"first_name" json:"first_name"`
	LastName          string             `bson:"last_name" json:"last_name"`
	EncryptedPassword string             `bson:"EncryptedPassword" json:"-"`
}

type UserUpdate struct {
	Email             string `bson:"email,omitempty" json:"email,omitempty"`
	FirstName         string `bson:"first_name,omitempty" json:"first_name,omitempty"`
	LastName          string `bson:"last_name,omitempty" json:"last_name,omitempty"`
	EncryptedPassword string `bson:"EncryptedPassword,omitempty" json:"-"`
}

type UserParams struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

type DeleteResult struct {
	DeleteCount int64 `bson:"n" json:"deleteCount"`
}
