package db

import (
	"context"
	"log"

	"github.com/juliosaraiva/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dburi    = "mongodb://localhost:27017"
	dbname   = "hotel-reservation"
	userColl = "users"
)

type UserStore interface {
	GetUserById(string) (*types.User, error)
	AddNewUser(*types.User) error
}

type MongoUserStore struct {
	client *mongo.Client
}

func (s *MongoUserStore) GetUserById(id int64) (*types.User, error) {
	ctx := context.Background()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	coll := client.Database(dbname).Collection(userColl)
	var user types.User
	if err = coll.FindOne(ctx, bson.M{"id": id}).Decode(&user); err != nil {
		log.Fatal(err)
	}

	return &user, nil
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		client: client,
	}
}

func (s *MongoUserStore) AddNewUser(username *types.User) error {
	ctx := context.Background()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}
	coll := client.Database(dbname).Collection(userColl)
	_, err = coll.InsertOne(ctx, &username)
	if err != nil {
		return err
	}

	return nil
}
