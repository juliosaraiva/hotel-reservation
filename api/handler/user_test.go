package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/juliosaraiva/hotel-reservation/api/db"
	"github.com/juliosaraiva/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dburi    = "mongodb://localhost:27017"
	dbname   = "test-hotel-reservation"
	userColl = "test-users"
)

type testdb struct {
	db.UserStore
}

func (tdb *testdb) TearDown(t *testing.T) {
	if err := tdb.UserStore.Drop(context.TODO()); err != nil {
		t.Fatal(t)
	}
}

func setup(t *testing.T) *testdb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	return &testdb{
		UserStore: &db.MongoUserStore{
			Collection: client.Database(dbname).Collection(userColl),
		},
	}
}

func TestPostUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.TearDown(t)

	app := fiber.New()
	userHandler := NewUserHandler(tdb.UserStore)
	app.Post("/api/v1/user", userHandler.CreateUser)

	params := types.UserParams{
		Email:     "bitcoin@btc.org",
		FirstName: "Satochi",
		LastName:  "Nakamoto",
		Password:  "my_bitcoins",
	}

	b, _ := json.Marshal(params)

	req := httptest.NewRequest("POST", "/api/v1/user", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	response, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}
	defer req.Body.Close()

	var user types.User
	json.NewDecoder(response.Body).Decode(&user)
	if user.FirstName != params.FirstName {
		t.Errorf("Expected FirstName %s but got %s", params.FirstName, user.FirstName)
	}

	if user.LastName != params.LastName {
		t.Errorf("Expected Lastname %s but got %s", params.LastName, user.LastName)
	}

	if user.Email != params.Email {
		t.Errorf("Expected Email %s but got %s", params.Email, user.Email)
	}

}
