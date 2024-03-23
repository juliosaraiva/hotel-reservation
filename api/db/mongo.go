package db

import (
	"context"

	"github.com/juliosaraiva/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStore interface {
	GetUserById(ctx context.Context, id string) (*types.User, error)
	GetUser(ctx context.Context) ([]*types.User, error)
	AddUser(ctx context.Context, user *types.User) (*types.User, error)
	UpdateUser(ctx context.Context, id string, params *types.UserUpdate) error
	DeleteUser(ctx context.Context, id string) (*mongo.DeleteResult, error)
}

type MongoUserStore struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func (s *MongoUserStore) GetUserById(ctx context.Context, id string) (*types.User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var user types.User
	if err := s.Collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *MongoUserStore) GetUser(ctx context.Context) ([]*types.User, error) {
	var users []*types.User
	cur, err := s.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cur.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil

}

func (s *MongoUserStore) AddUser(ctx context.Context, user *types.User) (*types.User, error) {
	insertedUser, err := s.Collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	user.ID = insertedUser.InsertedID.(primitive.ObjectID)

	return user, nil
}

func (s *MongoUserStore) UpdateUser(ctx context.Context, id string, params *types.UserUpdate) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	mapUser, err := bson.Marshal(bson.M{"$set": params})
	if err != nil {
		return err
	}
	_, err = s.Collection.UpdateOne(ctx, bson.M{"_id": oid}, mapUser)
	if err != nil {
		return err
	}

	return nil
}

func (s *MongoUserStore) DeleteUser(ctx context.Context, id string) (*mongo.DeleteResult, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	res, err := s.Collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, err
	}
	return res, nil
}
