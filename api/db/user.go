package db

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/juliosaraiva/hotel-reservation/internal/domain/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type user struct {
	Collection *mongo.Collection
}

func NewUser(collection *mongo.Collection) *user {
	return &user{
		Collection: collection,
	}
}

func (u *user) AddUser(ctx context.Context, user *models.User) (*models.User, error) {
	insertedUser, err := u.Collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	user.ID = insertedUser.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (u *user) GetUsers(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	cur, err := u.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cur.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (u *user) GetUserById(ctx context.Context, id string) (*models.User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user *models.User
	if err := u.Collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&user); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *user) UpdateUser(ctx context.Context, id string, params *models.UserUpdate) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	mapUser, err := bson.Marshal(bson.M{"$set": params})
	if err != nil {
		return err
	}

	_, err = u.Collection.UpdateOne(ctx, bson.M{"_id": oid}, mapUser)
	if err != nil {
		return err
	}
	return nil
}

func (u *user) DeleteUser(ctx context.Context, id string) (*models.DeleteResult, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	res, err := u.Collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, err
	}
	var deleteResult *models.DeleteResult
	resMarshal, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(resMarshal, deleteResult); err != nil {
		return nil, err
	}

	return deleteResult, nil
}

func (u *user) Drop(ctx context.Context) error {
	fmt.Println("--- Dropping user collection")
	return u.Collection.Drop(ctx)
}
