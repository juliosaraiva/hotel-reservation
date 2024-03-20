package db

import (
	"context"
	"log"

	"github.com/juliosaraiva/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserStore interface {
	GetUserById(id int64) (*types.User, error)
	AddNewUser(user *types.User) error
}

type MongoUserStore struct {
	Client     *mongo.Client
	DBUri      string
	DBName     string
	Collection string
}

func (s *MongoUserStore) GetUserById(id int64) (*types.User, error) {
	ctx := context.Background()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(s.DBUri))
	if err != nil {
		log.Fatal(err)
	}

	coll := client.Database(s.DBName).Collection(s.Collection)
	var user types.User
	if err = coll.FindOne(ctx, bson.M{"_id": id}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *MongoUserStore) AddNewUser(user *types.User) error {
	ctx := context.Background()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(s.DBUri))
	if err != nil {
		log.Fatal(err)
	}
	coll := client.Database(s.DBName).Collection(s.Collection)
	_, err = coll.InsertOne(ctx, &user)
	if err != nil {
		return err
	}

	return nil
}
