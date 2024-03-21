package db

import (
	"context"

	"github.com/juliosaraiva/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStore interface {
	GetUserById(ctx context.Context, id int64) (*types.User, error)
	AddNewUser(ctx context.Context, user *types.User) error
}

type MongoUserStore struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func (s *MongoUserStore) GetUserById(ctx context.Context, id int64) (*types.User, error) {
	var user types.User
	if err := s.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *MongoUserStore) AddNewUser(ctx context.Context, user *types.User) error {
	_, err := s.Collection.InsertOne(ctx, &user)
	if err != nil {
		return err
	}

	return nil
}
